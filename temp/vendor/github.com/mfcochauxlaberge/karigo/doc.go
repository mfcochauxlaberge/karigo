/*
Package karigo provides an HTTP API framework where the business logic is executed through transactions appended to an ordered log. The storage is seen as a list of keys and values and the transactions are simply a list of updates to apply to the keys.

An ordered log is an easy data structure to implement and mutate. In karigo, the log is meant to be exposed and handled directly. Components like search indexes and caches can simply read the log to stay up-to-date. The log can also be easily replicated through a consencus algorithm like Raft to improve avalaibility.
*/
package karigo
