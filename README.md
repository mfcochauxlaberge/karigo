# Karigo

<div align="center" style="text-align: center;">
  <img src="logo.png" height="120">
  <br>
  <a href="https://travis-ci.com/mfcochauxlaberge/karigo">
    <img src="https://travis-ci.com/mfcochauxlaberge/karigo.svg?branch=master">
  </a>
  <!-- <a href="https://codecov.io/gh/mfcochauxlaberge/karigo">
    <img src="https://codecov.io/gh/mfcochauxlaberge/karigo/branch/master/graph/badge.svg">
  </a> -->
  <a href="https://godoc.org/github.com/mfcochauxlaberge/karigo">
    <img src="https://godoc.org/github.com/golang/gddo?status.svg">
  </a>
</div>

Karigo is an API framework that follows the [JSON:API specification](https://jsonapi.org/format).

## State

This is a **work in progress**. Some of the features explained in this document might not exist yet. See the issues and the pull requests on GitHub to have an idea of what is being worked on.

See the [minimum viable project board](https://github.com/mfcochauxlaberge/karigo/projects/1) for a list of the first features being implemented and their current states. The work is being done on the `mvp` branch.

## Concepts

Karigo handles resources defined by their type. It accepts and returns valid JSON:API document.

### What is provided by the framework

 - Reading and writing valid JSON:API requests
 - Parse the URL (invluding its parameters)
 - Routing the URLs to the corresponding business logic
 - Provided default functions out-of-the-box if business logic is not necessary
 - Save each transaction in an exposed ordered log

### What the user has to do

 - Define the types
 - Provide validation rules
 - Write the business logic

### Type

A type has a name, attributes, and relationships.

 - Name
 - Attributes
   - String, number, boolean, etc
   - Can be nil
   - Have validation rules
 - Relationships
   - A string for to-one relationships
   - A slice of strings for to-many relationships
   - Can be empty
   - Can have an inverse relationship

### Resources

A resource is an instance of a type, like a row in a table.

During a GET requests, the parameters are parsed and used in order to return the appropriate collection with the correct filter, sorting, fields, include, etc.

### Log

Each request is appended to an ordered log. An entry in the log is simply a list of keys associated to some values. A key is a resource's field that will be set to the associated value.

### Transactions

Each request that modifies at least one resource needs to append a transaction to the log.

## Documentation

Documentation will be provided when the API is stable.

## Contributing

Contributions are **not** accepted at the moment.
