PROJECT_LEARNING.md
Personal Task Manager REST API - Development Journey
Purpose

This document records the resources, concepts, libraries, tools, and development decisions used while building the Personal Task Manager REST API.

It is intended for developers who want to understand how the project was developed rather than what the project does.

Development Approach

The project was not developed by writing all the code at once. Instead, it was built incrementally, where each layer was completed and tested before moving to the next.

Development order:

Designed project structure
Created Task model
Implemented Repository layer
Implemented Service layer
Implemented Handler layer
Registered HTTP routes
Tested APIs using Postman
Integrated MySQL
Debugged SQL and driver issues
Final API testing

This approach made debugging easier because every layer was verified before the next one was added.

Resources Used
Go Standard Library

The project intentionally uses Go's standard library instead of frameworks.

Packages used:

net/http

Purpose:

Create HTTP server
Register routes
Handle requests
Send responses

Used because it is lightweight and demonstrates how Go handles HTTP internally.

encoding/json

Purpose:

Decode request body
Encode response body

Example use:

POST request JSON → Go struct
Go struct → JSON response
database/sql

Purpose:

Provides a generic API for interacting with SQL databases.

Used for:

Opening database connections
Executing SQL statements
Querying rows
Scanning database results
errors

Purpose:

Create custom application errors.

Example:

Task not found
Invalid task ID
Empty title
strconv

Purpose:

Convert URL path parameters into integers.

Example:

/tasks/5

↓

5
strings

Purpose:

Used to extract task IDs from request paths.

time

Purpose:

Represents creation and update timestamps.

Used together with MySQL TIMESTAMP columns.

External Package
github.com/go-sql-driver/mysql

Purpose:

Connects Go applications with MySQL.

Reason for choosing:

Go's database/sql package defines database interfaces but does not communicate with MySQL by itself. A driver is required to implement that communication.

Important configuration:

parseTime=true

This allows MySQL DATETIME and TIMESTAMP values to be converted directly into Go's time.Time type.

Development Tools
Visual Studio Code

Used as the primary IDE.

Helpful features:

Go extension
Auto formatting
IntelliSense
Debugging
Integrated terminal
Go Toolchain

Commands used during development:

go run ./cmd

Runs the application.

go mod tidy

Downloads and cleans dependencies.

go get github.com/go-sql-driver/mysql

Installs the MySQL driver.

MySQL Workbench

Used to:

Create database
Create tables
Execute SQL queries
Verify inserted records
Debug SQL operations
Postman

Used for API testing.

Endpoints tested:

POST
GET
GET by ID
PUT
DELETE

Postman helped verify both HTTP responses and database changes.

Git

Used for version control.

Typical workflow:

git add .

git commit -m "Implemented MySQL integration"

git push
Important Concepts Learned
Layered Architecture

Separated responsibilities into:

Handler
Service
Repository
Database

Reason:

Improves readability and makes future changes easier.

Dependency Injection

Dependencies are passed through constructors instead of being created inside each layer.

Benefit:

Loose coupling
Easier testing
Better code organization
Repository Pattern

Database operations are isolated from business logic.

Benefit:

Changing the storage implementation requires modifications only in the repository layer.

REST API Design

Implemented standard CRUD endpoints following REST conventions.

SQL CRUD Operations

Practiced writing:

INSERT
SELECT
UPDATE
DELETE

queries directly from Go.

JSON Handling

Learned how Go converts:

JSON ⇄ Struct

using the encoding/json package.

Problems Encountered
Problem 1

Boolean value sent as a string.

Incorrect:

{
    "completed":"true"
}

Correct:

{
    "completed": true
}

Lesson:

JSON data types must match Go struct field types.

Problem 2

MySQL timestamp parsing error.

Error:

unsupported Scan, storing driver.Value type []uint8 into type *time.Time

Solution:

Added:

parseTime=true

to the MySQL DSN.

Lesson:

Database drivers sometimes require additional configuration to correctly map SQL types to Go types.

Problem 3

Repository migration.

Initially the repository stored data in memory.

After migration:

SQL queries replaced map operations.
Service layer remained largely unchanged because of proper separation of concerns.

Lesson:

A good architecture reduces the impact of implementation changes.

What This Project Demonstrates

This project demonstrates practical knowledge of:

Go backend development
REST API implementation
HTTP request lifecycle
Layered architecture
MySQL integration
SQL programming
JSON serialization
Error handling
Repository pattern
Dependency injection
API testing
Git workflow
References

Official Go Documentation:
https://go.dev/doc/

Go Packages:
https://pkg.go.dev/

Go by Example:
https://gobyexample.com/

MySQL Documentation:
https://dev.mysql.com/doc/

Go MySQL Driver:
https://github.com/go-sql-driver/mysql

Postman Documentation:
https://learning.postman.com/

Final Note

The objective of this project was not only to implement CRUD operations but also to understand how backend applications are structured in production-style projects. Each implementation decision—from using layered architecture to integrating MySQL—was made to improve maintainability, modularity, and scalability while reinforcing core backend development concepts.