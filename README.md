# Day 3 – JSON Encoding/Decoding and Validation in Go

## Overview

This project demonstrates how to build a simple REST API in Go that accepts JSON data, validates user input, and returns JSON responses. The API provides a POST endpoint to create a student record.

This project is designed as an internship-level exercise to understand JSON handling, request validation, and HTTP server development using Go's standard library.

---

## Objectives

- Learn JSON encoding and decoding
- Create a POST API endpoint
- Validate incoming request data
- Return structured JSON responses
- Understand HTTP status codes

---

## Project Structure

```
day3-json-api/
│── go.mod
│── main.go
│── model.go
│── handler.go
│── README.md
```

---

## Technologies Used

- Go 1.22+
- net/http
- encoding/json

No third-party packages are required.

---

## API Endpoint

### Create Student

**URL**

```
POST /student
```

**Content-Type**

```
application/json
```

---

## Request Body

```json
{
    "id": 1,
    "name": "Sneha",
    "email": "sneha@gmail.com",
    "age": 21
}
```

---

## Successful Response

**Status Code**

```
201 Created
```

**Response**

```json
{
    "success": true,
    "message": "Student added successfully",
    "data": {
        "id": 1,
        "name": "Sneha",
        "email": "sneha@gmail.com",
        "age": 21
    }
}
```

---

## Validation Rules

The API validates the following fields:

- Name cannot be empty.
- Email cannot be empty.
- Age must be greater than zero.

If validation fails, the API returns a **400 Bad Request** response.

Example:

```json
{
    "success": false,
    "errors": [
        "Name is required",
        "Email is required",
        "Age must be greater than zero"
    ]
}
```

---

## Invalid JSON Response

If the request body contains malformed JSON, the API returns:

```
400 Bad Request
```

Response:

```
Invalid JSON
```

---

## Running the Project

### Clone or Create the Project

```bash
mkdir day3-json-api
cd day3-json-api
```

Initialize Go module:

```bash
go mod init day3-json-api
```

Run the application:

```bash
go run .
```

Server Output:

```
Server running at http://localhost:8080
```

---

## Testing with cURL

### Valid Request

```bash
curl -X POST http://localhost:8080/student \
-H "Content-Type: application/json" \
-d '{
"id":1,
"name":"Sneha",
"email":"sneha@gmail.com",
"age":21
}'
```

---

### Validation Error

```bash
curl -X POST http://localhost:8080/student \
-H "Content-Type: application/json" \
-d '{
"id":2,
"name":"",
"email":"",
"age":0
}'
```

---

### Invalid JSON

```bash
curl -X POST http://localhost:8080/student \
-H "Content-Type: application/json" \
-d '{invalid}'
```

---

## Concepts Learned

- Go HTTP Server
- REST API Basics
- HTTP POST Method
- JSON Encoding
- JSON Decoding
- Struct Tags
- Input Validation
- Error Handling
- HTTP Status Codes
- JSON Responses

---

## HTTP Status Codes Used

| Status Code | Description |
|-------------|-------------|
| 201 | Student created successfully |
| 400 | Invalid JSON or validation failed |
| 405 | Method not allowed |

---

## Future Improvements

- Add GET endpoint to retrieve students.
- Store student data in a database.
- Add PUT and DELETE endpoints.
- Implement middleware for logging.
- Add unit tests.
- Add UUID-based student IDs.
- Improve email validation using regular expressions.

---

## Learning Outcome

By completing this project, you will gain hands-on experience with:

- Building REST APIs in Go
- Working with JSON data
- Request validation
- HTTP request and response handling
- Structs and methods
- Developing backend services using Go's standard library
