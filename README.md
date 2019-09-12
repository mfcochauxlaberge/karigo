# Karigo

<div align="center" style="text-align: center;">
  <img src="assets/logo.png" height="120">
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

Karigo is an API service that follows the [JSON:API specification](https://jsonapi.org/format).

Karigo is not a framework to build and run a binary. Karigo itself is the binary. When it is running for the first time, it always serves an empty API (no public schema and data). It then has to be customized using the CLI by adding types, validation, business logic, and more.

## State

This is a work in progress. It is not possible to make a production API with this tool yet.

## Use cases

### What Karigo can do

 - Read and write valid JSON:API requests
 - Parse the URL (including its parameters)
 - Route the endpoints to the corresponding business logic
 - Provide basic validation rules
 - Save each transaction in an exposed ordered log

### What the user has to do

 - Run the service (`karigo run`)
 - Define the types (names, attributes, and relationships)
 - Provide more specific validation rules
 - Write the business logic

## Concepts

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

### Resource

A resource is an instance of a type, like a row in a table.

During a GET requests, the parameters are parsed and used in order to return the appropriate collection with the correct filter, sorting, fields, include, etc.

### Log

The log is an ordered and append-only sequence of transactions.

Such a log makes a lot of tasks easier:

 - Replicate and synchronize the data
 - Replay the transactions to benchmark performance
 - Trigger events when a resource, field, or type is modified
 - Build a simple and powerful test suite

### Transaction

Each request modiying data results in a transaction appended to an ordered log. A transaction is a set of operations. An operation is the field of a resource and a value. Executing the operation means setting the resource's field to the value.

## Documentation

Documentation will be provided when the API is more stable.

## Contributing

Contributions are not accepted at the moment.
