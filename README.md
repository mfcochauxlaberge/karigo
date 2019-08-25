# karigo

<div align="center" style="text-align: center;">
  <img src="logo.png" height="120">
  <br>
  <a href="https://travis-ci.com/mfcochauxlaberge/karigo">
    <img src="https://travis-ci.com/mfcochauxlaberge/karigo.svg?branch=mvp">
  </a>
  <!-- <a href="https://codecov.io/gh/mfcochauxlaberge/karigo">
    <img src="https://codecov.io/gh/mfcochauxlaberge/karigo/branch/master/graph/badge.svg">
  </a> -->
  <a href="https://godoc.org/github.com/mfcochauxlaberge/karigo">
    <img src="https://godoc.org/github.com/golang/gddo?status.svg">
  </a>
</div>

karigo aims to be an HTTP API framework where the business logic is executed through transactions appended to an ordered log. The storage is seen as a list of keys and values and the transactions are simply a list of updates to apply to the keys.

An ordered log is an easy data structure to implement and mutate. In karigo, the log is meant to be exposed and handled directly. Components like search indexes and caches can simply read the log to stay up-to-date. The log can also be easily replicated through a consencus algorithm like Raft to improve avalaibility.

## State

This is a work in progress.

See the [minimum viable project board](https://github.com/mfcochauxlaberge/karigo/projects/1) for a list of the first features being implemented and their current states. The work is being done on the `mvp` branch.

## Documentation

There is no documentation so far, except for the source code.

## Contributing

Contributions are not accepted for now.