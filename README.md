# Task Management API (GoLang)

This is a GoLang API for managing tasks. It provides functionalities for registering users, logging in, creating, retrieving, updating, and deleting tasks.

## Features

* User Registration and Login
* Task Creation, Retrieval (Single and All), Update, and Deletion
* JWT Authentication (for secure access)

## Environment Variables

| Variable                  | Description                                           |
|---------------------------|-------------------------------------------------------|
| DB_NAME                   | ashutosh                                              |
| DB_PASSWORD               | rtx2080ti                                             |
| JWT_KEY                   |  task-management-jwt-key                              |
| USER_COLLECTION           |  task                                                 |
| TASK_COLLECTION           |  user                                                 |
-------------------------------------------------------------------------------------
## API Documentation

**Base URL:** 
`http://localhost:5000/users`

**Authentication:**

This API uses JWT authentication. Users can register and obtain access and refresh tokens upon successful login. These tokens are required in the Authorization header for subsequent requests.

**Endpoints:**

**1. Register User:**

**Method:** POST
**Endpoint:** `/register`

**Request Body:**

```json
{
  "firstname": "string",
  "lastname": "string",
  "email": "string",
  "password": "string"
}

Response:
{
  "accessToken": "string",  // JWT access token
  "expiresIn": number,      // Access token expiration time in seconds
  "refreshToken": "string", // JWT refresh token
  "refreshTokenExpiresIn": number // Refresh token expiration time in seconds
}
```

**2. Login:**

**Method:** POST
**Endpoint:** /login

```json
Request Body:

JSON
{
  "email": "string",
  "password": "string"
}

Response:
{
  "accessToken": "string",  // JWT access token
  "expiresIn": number,      // Access token expiration time in seconds
  "refreshToken": "string", // JWT refresh token
  "refreshTokenExpiresIn": number // Refresh token expiration time in seconds
}
```

**3.Create Task:**

**Method:** POST
**Endpoint:** /tasks/
Authorization: Required (Bearer token)

Request Body:
```json
{
  "task": "string" // The description of the task
}

Response:

{
  "message": "Task created"
}
```

**4.Get Single Task:**

**Method:** GET
**Endpoint:** /tasks/{taskId}

Authorization: Required (Bearer token)

Path Parameters:
{taskId} (string): The ID of the task to retrieve.

```json
Response:
{
  "_id": "string",
  "user_id": "string",
  "task": "string",
  "created_at": "string" // ISO 8601 formatted date and time
}
```

**5. Get All Tasks:**

**Method:** GET
**Endpoint:** /tasks/

Authorization: Required (Bearer token)

Response:
```json
[
  {
    "_id": "string",
    "user_id": "string",
    "task": "string",
    "created_at": "string" // ISO 8601 formatted date and time
  },
  // ... (other tasks)
]
```

**6. Update Task:**

**Method: PUT**
**Endpoint: /tasks/{taskId}**

Authorization: Required (Bearer token)

Path Parameters:
{taskId} (string): The ID of the task to update.
Request Body:

```json
{
  "task": "string" // The updated description of the task
}

Response:
{
  "message": "Task updated"
}
```


**7. Delete Task:**

**Method: DELETE**
**Endpoint: /tasks/{taskId}**

Authorization: Required (Bearer token)

Path Parameters:
{taskId} (string): The ID of the task to delete.

```json
Response:
{
  "message": "Task deleted"
}
```

**SAMPLE SEVER EXECUTION RESPONSE**
***server connection successful***

 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.52.5                   │ 
 │               http://127.0.0.1:5000               │ 
 │       (bound on host 0.0.0.0 and port 5000)       │ 
 │                                                   │ 
 │ Handlers ............ 13  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 28240 │ 
 └───────────────────────────────────────────────────┘ 

