Golang -- Envornment veriables -- need to add these before executing the project

*******************************
DB_NAME=ashutosh                  
DB_PASSWORD=rtx2080ti             
JWT_KEY=task-management-jwt-key   
USER_COLLECTION=user              
TASK_COLLECTION=task            
******************************

API DOC
*******

BASE-URL === localhost:5000/users -- BY DEFAULT THE LOCAL HOST IS SET TO 5000

Register User
-------------

Method: POST
Endpoint: /register
Request Body:
{
  "firstname": "ashutosh",
  "lastname": "pandey",
  "email": "adsdfsarsh@gmail.com",
  "password": "password123"
}

Response:
Status Code: 200 OK
BODY:
{
    "accessToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkzMzY1NDksImlhdCI6MTcyOTI5MzM0OSwidHlwZSI6ImFjY2VzcyIsInVpZCI6IjY3MTJlYzI1OGEyMGUzMTVkY2UyMTNiMiIsImVtYWlsIjoiIn0.VlU_PUQnLS5yaaXiThEO6Q2-FaUYJtgrYld6sRcMwWL50Cs5xgw4O2SWMIV9i89FruPEMX7V2a7lHqcZGTiHhw",
    "expiresIn": 43200,
    "refreshToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE4ODUzNDksImlhdCI6MTcyOTI5MzM0OSwidHlwZSI6InJlZnJlc2giLCJ1aWQiOiI2NzEyZWMyNThhMjBlMzE1ZGNlMjEzYjIifQ.NNbA7lo5cnG4QampucYCup7eTyBwwfLI-PFGFR_szyLDofY_lbowYNbzGur091nYwdDzVWLZx-4qkDy1b6ra6w",
    "refreshTokenExpiresIn": 2592000
}


Login
-----
Method: POST
Endpoint: /login
Request Body:
{
  "email": "adsdfsarsh@gmail.com",
  "password": "password123"
}

Response:
Status Code: 200 OK
BODY:
{
    "accessToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkzMzY1NDksImlhdCI6MTcyOTI5MzM0OSwidHlwZSI6ImFjY2VzcyIsInVpZCI6IjY3MTJlYzI1OGEyMGUzMTVkY2UyMTNiMiIsImVtYWlsIjoiIn0.VlU_PUQnLS5yaaXiThEO6Q2-FaUYJtgrYld6sRcMwWL50Cs5xgw4O2SWMIV9i89FruPEMX7V2a7lHqcZGTiHhw",
    "expiresIn": 43200,
    "refreshToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE4ODUzNDksImlhdCI6MTcyOTI5MzM0OSwidHlwZSI6InJlZnJlc2giLCJ1aWQiOiI2NzEyZWMyNThhMjBlMzE1ZGNlMjEzYjIifQ.NNbA7lo5cnG4QampucYCup7eTyBwwfLI-PFGFR_szyLDofY_lbowYNbzGur091nYwdDzVWLZx-4qkDy1b6ra6w",
    "refreshTokenExpiresIn": 2592000
}


Create Task
--------------
Method: POST
Endpoint: /tasks/
Request Body:
{
    "task": "SOME TASK TO CREATE"  // The description of the task
}

Response:
Status Code: 200 OK
Body:
{
    "message": "Task created"
}

Get Single Task
---------------
Method: GET
Endpoint: /tasks/{taskId}
Path Parameters:
taskId (string): The ID of the task to retrieve.


Response:
Status Code: 200 OK
Body:
{
    "_id": "6712e0405bc27c94aa1ec906",
    "user_id": "6712c18bf68bd665883b4c8c",
    "task": "new task api",
    "created_at": "2024-10-18T22:25:04.212Z"
}

Get All Tasks
-------------
Method: GET
Endpoint: /tasks/

Response:
Status Code: 200 OK
[ 
   .....
    {
        "_id": "6712e0405bc27c94aa1ec906",
        "user_id": "6712c18bf68bd665883b4c8c",
        "task": "new task api",
        "created_at": "2024-10-18T22:25:04.212Z"
    },
    {
        "_id": "6712e0405bc27c94aa1ec906",
        "user_id": "6712c18bf68bd665883b4c8c",
        "task": "new task api",
        "created_at": "2024-10-18T22:25:04.212Z"
    }
    ......
   
]



Update Task
-----------

Method: PUT
Endpoint: /tasks/{taskId}
Path Parameters:
taskId (string): The ID of the task to update.
Request Body:
Copy code
{
    "task": "SOME GIVEN TASK"  // The updated description of the task
}

Response:
Status Code: 200 OK
Body:
{
    "message": "Task updated"
}

Delete Task
-----------
Method: DELETE
Endpoint: /tasks/{taskId}
Path Parameters:
taskId (string): The ID of the task to delete.

Response:
Status Code: 200 OK
Body:
{
    "message": "Task deleted"
}



