# go-user-microservice-assessment
A simple Go-based microservice for managing users with CRUD operations and PostgreSQL integration.

---
## Table of Contents

1. [Project Structure](#project-structure)
2. [Install dependencies](#install-dependencies)
3. [Set up the environment](#set-up-the-environment)
4. [Running Migrations](#running-migrations)
5. [Start the Server](#start-the-server)
6. [Testing the API](#testing-the-api)

---
### Project Structure
```
golang-microservice/
│── main.go
│── config/
│   └── config.go
│── models/
│   └── user.go
│── routes/
│   └── user_routes.go
│── controllers/
│   └── user_controller.go
│── migrations/
│   └── migrate.go
│── .env
│── go.mod
│── go.sum
│── README.md
```

---
### Install dependencies
```
   go mod tidy
```
---
### Set up the environment
```
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=userdb
   DB_PORT=5432
   SERVER_PORT=8080
```
---
### Running Migrations
```
   go run migrations/migrate.go
```
---
### Start the Server
```
   go run main.go
```
---
### Testing the API

#### Create a User
```
   POST /users
   {
      "name": "Jordan Lee",
      "email": "jordanlee19930130@gmail.com"
   }
```
![createUser](./result/createUsers.png)

#### Get User by ID
```
   GET /users/{id}
```
![getUserValid](./result/getUserValid.png)
#### Update User
```
   PUT /users/{id}
   {
   "name": "John Updated",
   "email": "john.updated@example.com"
   }
```
![updateUser](./result/updateUserValid.png)
#### Delete User
```
   DELETE /users/{id}
```

