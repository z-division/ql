package state_machine

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
    // Who owns the state machine (e.g., host name, or better yet logical
    // replication group name, etc.)
    OwnerId string

    // The state machine's name (e.g., "database", "txn-XXXXXX", etc.)
    Name string

    // Zero represents the state machine's initial state (i.e., All WAL
    // entries should have positive numbers)
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
// TODO(patrick): add checkpointing
// TODO(patrick): add compaction
// TODO(patrick): add schema support
// TODO(patrick): add sharded database specific operations, e.g., range split
// (self-reminder: range split works best using pessimistic concurrency control)
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
    CreateTransaction(
        dbLsn LSN,
        clientLsn LSN,
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
}

// XXX(patrick): Maybe expose snapshotId to enable users to read at
// different snapshot points.
type Transaction interface {
    // State transition methods

    // Compact the values local to the transaction and perform various
    // preflight checks.  It's safe to finalize multiple times.  ExposeUpdates,
    // Put, and Delete are no longer valid action after finalizing.
    Finalize(txnLsn LSN) error

    // Create a new snapshot which exposes the transaction's pending updates,
    // e.g., after each sql modification statement.
    ExposeUpdates(txnLsn LSN) error

    Put(txnLsn LSN, key string, value string) error

    Delete(txnLsn, key string) error

    // This is used for implementing serializable OCC transactions, to "grab
    // read locks".
    //
    // Pessimistic transactions and repeatable reads OCC transactions do not
    // need to call this.
    ReadIntent(txnLsn, rng Range) error

    // Idempotent read-only methods
    Get(minExpectedTxnLsn LSN, key string) (ResultSet, error)

    // XXX(patrick): maybe this is a read option rather than an api?
    Has(minExpectedTxnLsn LSN, key string) (ResultSet, error)

    Scan(
        minExpectedTxnLsn LSN,
        rng Range,
        continuation ScanContinutationToken) (
        ResultSet,
        ScanContinutationToken,
        error)
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

// TODO(patrick): figure out the details
type TransactionOptions struct {
    Pessimistic bool

    // pessimistic transactions must predefine the read/write range
    ReadRange Range
    UpdateRange Range
}

// TODO(patrick): figure out the details
type Range struct {
}

// TODO(patrick): figure out the details
type ScanContinutationToken struct {
}

// TODO(patrick): figure out the details
type ResultSet struct {
}
