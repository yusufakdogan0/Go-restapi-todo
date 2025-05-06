
📝 TO-DO App API (Go + Gin + JWT)

This is a mock back-end REST API for managing to-do lists and their steps. Authentication is handled using JWT tokens. No real database is used — all data is stored in memory.

The deployment is done in railway. The base URL will be `https://go-restapi-todo-production.up.railway.app` for remote access

The base in localhost will be `http://localhost:8080` if no configuration is done.

---

👤 Registered Users

There are 3 predefined users:

| Username | Password   | Role  |
|----------|------------|-------|
| admin    | admin123   | admin |
| user1    | user123    | user  |
| user2    | user456    | user  |

- `admin` can see **all** users' to-do lists.
- `user` can only see and manipulate **their own** data.

---

🔐 Authentication

POST /login

Request:

`{
  "username": "user1",
  "password": "user123"
}`

Response:

`{
  "token": "JWT_TOKEN_STRING"
}`

Use this token in the Authorization header for all protected endpoints:

Authorization: Bearer YOUR_TOKEN

---

📋 Endpoints

POST /todo-lists
Create a new to-do list for the authenticated user.

Body:

`{
  "name": "My First List"
}`

Response:
Returns the created list object with ID and timestamps.

---

GET /todo-lists
Fetch to-do lists:

- If the user is an admin, returns all users' lists.
- If the user is a normal user, returns only their own lists.
- Only non-deleted lists are shown.
- Each list includes completion_percentage calculated from its steps.

Response:

`[
  {
    "id": "uuid",
    "name": "My First List",
    "username": "user1",
    "completion_percentage": 66.67,
    ... 
  }
]
`

---

DELETE /todo-lists/:id
Soft-delete a to-do list by setting its DeletedAt field. This also hides it from future fetches.

Response:

`{
  "message": "List soft-deleted"
}`

---

POST /todo-steps
Add a new step to a to-do list.

Body:

`{
  "todo_list_id": "list-uuid",
  "content": "Go Project"
}`

Response:
Returns the created step object.

---

PUT /todo-steps/:id
Update the content or completion status of a to-do step.

Body:

`{
  "content": "Go Project",
  "is_done": true
}`

Response:
Returns the updated step object.

---

DELETE /todo-steps/:id
Soft-delete a to-do step by setting its DeletedAt field.

Response:

`{
  "message": "Step soft-deleted"
}
`

