# Personal Task Manager REST API

A backend application developed in Go that provides RESTful APIs for managing personal tasks. The application follows a layered architecture consisting of Handler, Service, Repository, and MySQL Database layers.

The project also includes an interactive Console Client that communicates with the REST API over HTTP, allowing users to test and manage tasks without using Postman.

---

# Project Overview

The Personal Task Manager is a CRUD-based REST API application designed to manage daily tasks efficiently.

The project demonstrates backend development concepts such as:

- REST API Development
- Layered Architecture
- MySQL Database Integration
- HTTP Client and Server Communication
- JSON Encoding & Decoding
- Error Handling
- Go Project Structure

---

# Features

- Create Task
- List All Tasks
- Get Task By ID
- Update Task
- Delete Task
- Mark Task Completed
- Interactive Console Client
- MySQL Database Integration
- RESTful API Design
- Layered Architecture

---

# Project Architecture

```
                User
                  │
                  ▼
      Console Client / Postman
                  │
                  ▼
             HTTP Server
                  │
                  ▼
          Handler Layer
                  │
                  ▼
          Service Layer
                  │
                  ▼
        Repository Layer
                  │
                  ▼
             MySQL Database
```

---

# Layer Responsibilities

## Handler Layer

- Receives HTTP requests
- Parses request data
- Calls the Service layer
- Sends JSON responses

---

## Service Layer

- Implements business logic
- Performs input validation
- Coordinates operations
- Returns processed results

---

## Repository Layer

- Executes SQL queries
- Performs CRUD operations
- Interacts directly with MySQL
- Maps database records into Go structs

---

## Database Layer

Stores application data permanently.

---

# Project Structure

```
Personal-Task-Manager-API/

├── cmd
│   ├── server
│   │   └── main.go
│   │
│   └── console
│       ├── main.go
│       ├── client.go
│       ├── utils.go
│       ├── task.go
│       ├── create_task.go
│       ├── list_tasks.go
│       ├── get_task.go
│       ├── update_task.go
│       ├── delete_task.go
│       └── complete_task.go
│
├── docs
│   └── PROJECT_LEARNING.md
│
├── internal
│   ├── database
│   │   └── mysql.go
│   │
│   ├── handlers
│   │   └── task_handler.go
│   │
│   ├── models
│   │   └── task.go
│   │
│   ├── repository
│   │   └── task_repository.go
│   │
│   └── service
│       └── task_service.go
│
├── README.md
├── go.mod
├── go.sum
└── .gitignore
```

---

# Technologies Used

| Technology | Purpose |
|------------|---------|
| Go | Backend Development |
| MySQL | Database |
| net/http | REST API Server |
| database/sql | Database Access |
| go-sql-driver/mysql | MySQL Driver |
| JSON | API Communication |
| Git | Version Control |
| GitHub | Repository Hosting |

---

# Installation

Clone the repository

```bash
git clone https://github.com/Ajay-163/Personal-Task-Manager-API.git
```

Move into the project

```bash
cd Personal-Task-Manager-API
```

Install dependencies

```bash
go mod tidy
```

---

# Database Setup

Create a MySQL database.

```sql
CREATE DATABASE task_manager;
```

Create the Tasks table.

```sql
CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP
);
```

Update the database credentials in

```
internal/database/mysql.go
```

---

# Running the REST API

Start the server.

```bash
go run ./cmd/server
```

The server runs on

```
http://localhost:8080
```

---

# Running the Console Client

Open another terminal.

Run

```bash
go run ./cmd/console
```

The console provides an interactive menu.

```
====================================
     PERSONAL TASK MANAGER
====================================

1. Create Task
2. List Tasks
3. Get Task By ID
4. Update Task
5. Delete Task
6. Mark Task Completed
7. Exit
```

The Console Client communicates with the REST API using HTTP requests and provides an alternative to testing with Postman.

---

# REST API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | /tasks | Create Task |
| GET | /tasks | Get All Tasks |
| GET | /tasks/{id} | Get Task By ID |
| PUT | /tasks/{id} | Update Task |
| DELETE | /tasks/{id} | Delete Task |

---

# Request Example

## Create Task

POST

```
/tasks
```

Request Body

```json
{
    "title":"Learn Go",
    "description":"Study REST APIs",
    "completed":false
}
```

Response

```json
{
    "id":1,
    "title":"Learn Go",
    "description":"Study REST APIs",
    "completed":false
}
```

---

# Console Workflow

```
Start Server

↓

Start Console

↓

Create Task

↓

View Tasks

↓

Update Task

↓

Mark Completed

↓

Delete Task
```

---

# Learning Outcomes

This project helped in understanding:

- REST API Development
- Layered Architecture
- CRUD Operations
- HTTP Request Handling
- HTTP Client Development
- JSON Encoding & Decoding
- MySQL Integration
- Repository Pattern
- Error Handling
- Go Modules
- Git & GitHub Workflow

---

# Author

**Ajay Bondi**

Backend Developer

Go |MySQL
