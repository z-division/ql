package state_machine

type StateMachineId struct {
	// Which type of state machine this is.
	MachineType string

	// Who owns the state machine (e.g., host name, or better yet logical
	// replication group name, etc.)
	OwnerId string

	// The state machine's name (e.g., "database", "txn-XXXXXX", etc.)
	Name string
}

// "Log sequence number".  Unlike traditional lsn, this includes the owner id
// and state machine name since each owner may host multiple state machines
// and state machines may interact each other (e.g., the database is a
// long-lived state machine, each local transaction is also a short-lived state
// machine, and a distributed transaction's 2PC coordinator is also a
// short-lived state machine).  LSNs are total order only if the LSNs'
// (OwnerId, StateMachineName) are the same.  LSNs are partially ordered
// otherwise.
//
// XXX(patrick): maybe optimize this if it's too bloated
type LSN struct {
	StateMachineId

	// Zero represents the state machine's initial state.
	SequenceNumber uint64
}

// XXX(patrick): maybe expose more generic StateMachine methods, e.g.,
//  BeginCheckpoint() error
//  FinishCheckpoint() (Checkpoint, error)
type StateMachine interface {
	// TODO(patrick): figure out exactly API
	// Apply(LSN, eventRecord) error

	CurrentLSN() LSN
}

// A passive state machine which implements a low level homogenous relational
// database (aka a table).
//
// TODO(patrick): add sharded database specific operations, e.g., range split
// (self-reminder: range split works best using pessimistic concurrency control)
//
// For simplicity, the database may have at most two active schema running
// concurrently (only during alteration).
//
// self-reminder: schema alternation may cause record ordering to change across
// levels.  We have to be careful with how the tree is scanned / compacted.
// For compaction, rather than resorting an entire level, we can cheat and
// rewrite the level as mutliple levels with non-overlapping keys (and sort
// piecemeal)
//
// XXX(patrick): Does the checkpointing API make sense?
type Database interface {
	StateMachine

	// State transition methods

	// A transaction is uniquely identified by the clientLsn.
	//
	// Depending on the transaction's concurrency control mode, a newly created
	// transaction may or may not be immediately available for use.  In
	// particular, a pessimisitic transaction is available only when it
	// controls the entire requested key range.  See GetTransaction for
	// additional details.
	//
	// CreateTransaction is idempotent relative to both the database itself
	// the client's clientLsn (Client idempotency is needed for 2PC coordinator
	// recovery).
	//
	// If the clientLsn is not specified, the dbLsn is used as the clientLsn.
	// XXX(patrick): should we have an explicit method to generate the
	// transaction id?
	//
	// If the client lsn is smaller than the latest known lsn for that client,
	// this does nothing and return an invalid argument error.
	//
	// If the schemaVersion does not match the serving schema version, this
	// returns an error.
	CreateTransaction(
		dbLsn LSN,
		clientLsn LSN,
		schemaVersion SchemaVersion,
		options TransactionOptions) error

	// Prepare the transaction for commit.
	//
	// If the transaction is in optimistic concurrency control mode, this
	// performs a final check to ensure the transaction does not conflicit
	// with other transactions.  (The check can be skipped if the transaction
	// is in pessimistic concurrency control mode).
	//
	// NOTE(patrick): The transaction must be finalized before entering
	// the prepare state.
	//
	// NOTE(patrick): PrepareTransaction and CommitTransaction are implemented
	// as two distinct state transition in order to accommodate 2PC-style
	// distributed transactions.
	PrepareTransaction(dbLsn LSN, clientLsn LSN) error

	// NOTE(patrick): All transacation must be explicitly committed or
	// aborted.

	// Commit a prepared transaction and free up resources.
	Commit(dbLsn LSN, clientLsn LSN) error

	// Abort a transaction and free up resources.  A transaction may be
	// aborted at any time during its lifespan.
	Abort(dbLsn LSN, clientLsn LSN) error

	// Signal to the database that it should aler the tree using the new
	// schema.  The alternation are done as part of compaction.
	//
	// NOTE(patrick): The serving schema remains unmodified until
	// FinishSchemaAlteration is called.
	BeginSchemaAlteration(dbLsn LSN, newSchema Schema) (SchemaVersion, error)

	// Use new schema as the serving schema.
	// If GetSchema's servingLevelCount is non-zero, this returns an error.
	FinishSchemaAlteration(dbLsn LSN) error

	Compact(
		dbLsn LSN,
		candidate CompactionCandidate) (
		ContinuationToken,
		error)

	// BeginCheckpoint in the background.  The checkpoint is not finalized
	// until FinishCheckpoint is called.
	//
	// checkpoints don't overlap.  This returns an error if checkpointing
	// is already in progress.
	//
	// The suffix is useful for attaching timestamp info to checkpoint, etc.
	BeginCheckpoint(dbLsn LSN, checkpointSuffix string) error

	// Finish checkpointing and make the new checkpoint available via
	// ListCheckpoints/GetCheckpoint.
	//
	// This is a blocking operation.  IsCheckpointReady should
	// return true before calling this.
	FinishCheckpoint(dbLsn LSN) error

	// Remove a checkpoint and free up resources.
	RemoveCheckpoint(dbsLsn LSN, checkpointId CheckpointId) error

	// Idempotent read-only methods

	// If the transaction is not ready (in pessimistic mode), this returns a
	// not ready error.
	//
	// If the transaction is not found in database, this returns a not found
	// error.
	GetTransaction(
		minExpectedDbLsn LSN,
		clientLsn LSN) (
		Transaction,
		error)

	GetSchema(minExpectedDbLsn LSN) (
		serving Schema,
		servingLevelCount int,
		alteringTo Schema, // nil if schema is not changing
		alteringToLevelCount int,
		err error)

	GetCompactionCandidates(minExpectedDbLsn LSN) (
		[]CompactionCandidate,
		error)

	IsCheckpointReady(minExpectedLsn LSN) (bool, error)

	ListCheckpoints(minExpectedLsn LSN) (
		[]CheckpointId, // newest to oldest
		error)

	GetCheckpoint(
		minExpectedLsn LSN,
		checkpointId CheckpointId) (
		Checkpoint,
		error)
}

type CompactionCandidate struct {
	Level    int
	InMemory bool
	ContinuationToken
}

// Transaction workflow:
//    # pessimistic transactions must specify all the keys upfront
//    db.CreateTransaction(...)
//
//    for true {
//        txn, err = db.GetTransaction(...)
//        if err != nil {
//            if err == NotReady {  # pessimistic transaction not ready
//                wait
//                continue
//            }
//            return err
//        }
//        break
//    }
//
//    ...
//    do stuff with txn (in pessimistic model, this errors if
//    accessing/modifying data outside of range specified in CreateTransaction)
//    ...
//
//    txn.Finalize()
//    db.Prepare(...)
//    db.Commit(...)  # or Abort
type Transaction interface {
	// State transition methods

	// Compact the values local to the transaction and perform various
	// preflight checks.  It's safe to finalize multiple times.  ExposeUpdates,
	// Put, and Delete are no longer valid action after finalizing.
	Finalize(txnLsn LSN) error

	// Create a new snapshot which exposes the transaction's pending updates,
	// e.g., after each sql modification statement.
	ExposeUpdates(txnLsn LSN) error

	// NOTE(patrick): Put conforms to the serving schema.  Tuple fields are
	// positionally ordered the same as the schema.
	Put(txnLsn LSN, fullRecord Tuple) error

	// NOTE(patrick): Delete conforms to the serving schema.  Tuple fields are
	// positionally ordered the same as the schema.
	Delete(txnLsn, key Tuple) error

	// This is used for implementing serializable OCC transactions, to "grab
	// read locks".
	//
	// Pessimistic transactions and repeatable reads OCC transactions do not
	// need to call this.
	ReadIntent(txnLsn, rng Range) error

	// Idempotent read-only methods

	// Tuple fields are positionally ordered the same as the specified
	// projection.
	//
	// NOTE(patrick): Has check is just Get with no selected column.
	Get(
		minExpectedTxnLsn LSN,
		keys Tuples,
		projection *Projection) (
		Tuples,
		error)

	// Tuple fields are positionally ordered the same as the specified
	// projection.
	Scan(
		minExpectedTxnLsn LSN,
		rng Range,
		projection *Projection,
		batchSize int,
		continuation ContinuationToken) (
		Tuples,
		ContinuationToken,
		error)
}

// TODO(patrick): figure out the details
type TransactionOptions struct {
	Pessimistic bool

	// pessimistic transactions must predefine the read/write range
	ReadRange   Range
	UpdateRange Range
}

// TODO(patrick): support range union
type Range struct {
	Start Tuple // Unbounded if nil.  Inclusive if specified.
	End   Tuple // Unbounded if nil.  Inclusive if specified.
}

// NOTE(patrick): nil projection means project all columns
type Projection struct {
	ColumnNames []string
}

type ContinuationToken struct {
	RootLevelId string

	Depth int // number of levels to include.  non-positive means entire tree.

	nextKey Tuple // inclusive
}

// TODO(patrick): figure out the details
type Schema struct {
}

// TODO(patrick): figure out the details
type SchemaVersion struct {
}

// We assume that tuple fields are positionally ordered the same as the schema
// (or the column projection)
type Tuple interface {
	NumColumns() int

	Get(col int) (interface{}, error)
}

type Tuples interface {
	NumTuples() int

	// return the i-th tuple
	Get(i int) (Tuple, error)
}

type CheckpointId string

// TODO(patrick): figure out the details.  Maybe expose io reader for
// a tar format or something
type Checkpoint struct {
}
