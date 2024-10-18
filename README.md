# Task Management API (GoLang)

This is a GoLang API for managing tasks. It provides functionalities for registering users, logging in, creating, retrieving, updating, and deleting tasks.

## Features

* User Registration and Login
* Task Creation, Retrieval (Single and All), Update, and Deletion
* JWT Authentication (for secure access)

## Environment Variables

| Variable                 | Description                                            |
|---------------------------|-------------------------------------------------------|
| DB_NAME                   | ashutosh                                              |
| DB_PASSWORD               | rtx2080ti                                             |
| JWT_KEY                   |  task-management-jwt-key                              |
| USER_COLLECTION           |  task                                                 |
| TASK_COLLECTION           |  user                                                 |

## API Documentation

**Base URL:** 
`http://localhost:5000`

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

JSON
{
  "accessToken": "string",  // JWT access token
  "expiresIn": number,      // Access token expiration time in seconds
  "refreshToken": "string", // JWT refresh token
  "refreshTokenExpiresIn": number // Refresh token expiration time in seconds
}
Use code with caution.

2. Login:

Method: POST
Endpoint: /login

Request Body:

JSON
{
  "email": "string",
  "password": "string"
}
Use code with caution.

Response:

JSON
{
  "accessToken": "string",  // JWT access token
  "expiresIn": number,      // Access token expiration time in seconds
  "refreshToken": "string", // JWT refresh token
  "refreshTokenExpiresIn": number // Refresh token expiration time in seconds
}
Use code with caution.

3. Create Task:

Method: POST
Endpoint: /tasks/

Authorization: Required (Bearer token)

Request Body:

JSON
{
  "task": "string" // The description of the task
}
Use code with caution.

Response:

JSON
{
  "message": "Task created"
}
Use code with caution.

4. Get Single Task:

Method: GET
Endpoint: /tasks/{taskId}

Authorization: Required (Bearer token)

Path Parameters:

{taskId} (string): The ID of the task to retrieve.
Response:

JSON
{
  "_id": "string",
  "user_id": "string",
  "task": "string",
  "created_at": "string" // ISO 8601 formatted date and time
}
Use code with caution.

5. Get All Tasks:

Method: GET
Endpoint: /tasks/

Authorization: Required (Bearer token)

Response:

JSON
[
  {
    "_id": "string",
    "user_id": "string",
    "task": "string",
    "created_at": "string" // ISO 8601 formatted date and time
  },
  // ... (other tasks)
]
Use code with caution.

6. Update Task:

Method: PUT
Endpoint: /tasks/{taskId}

Authorization: Required (Bearer token)

Path Parameters:

{taskId} (string): The ID of the task to update.
Request Body:

JSON
{
  "task": "string" // The updated description of the task
}
Use code with caution.

Response:

JSON
{
  "message": "Task updated"
}
Use code with caution.

7. Delete Task:

Method: DELETE
Endpoint: /tasks/{taskId}

Authorization: Required (Bearer token)

Path Parameters:

{taskId} (string): The ID of the task to delete.
Response:

JSON
{
  "message": "Task deleted"
}
Use code with caution.

Installation
Clone the repository:
git clone [invalid URL removed]
Change to the project directory:
cd task-management-api
Install dependencies:
go get
Create a .env file and add the following environment variables:
DB_NAME
DB_PASSWORD
JWT_KEY
USER_COLLECTION
TASK_COLLECTION
Run the server:
go run main.go
The API should now be running on port 5000. You can access the API endpoints using a web browser or any HTTP client.
