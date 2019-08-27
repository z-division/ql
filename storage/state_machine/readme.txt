State machine design self-reminder:
- the state machine is fully passive.  All changes to the state machine are
  triggered by external clients or monitors
- state machine is drive by an external WAL.  The state machine must handle log
  replay idempotently:
    - state mutation methods must take a log sequence number as argument.  If
      the input log sequence number is less than the latest known LSN, the
      state machine must remain unmodified and return an error
      (TODO(patrick): define error)
- no state machine method can block indefinitely.
- whenever possible, state transitions should be separated from idempotent
  reads, and should not return values.
- whenever possible, large blocking operations should be split into
  multiple state machine transcations.
- different state machiens may or may not share the same log (logs are
  completely ordered within a state machine, but may be partially ordered
  across different state machines)
- the state machine transition should by time/timing agnostic

- read-only method should take a minimum expected lsn as an argument (in case
  the state machine is lagging)
  - maybe this can be independent of the state machine implementation?
