# go-user-microservice-assessment
A simple Go-based microservice for managing users with CRUD operations and PostgreSQL integration.

---
## Table of Contents

1. [Project Structure](#project-structure)
2. [Setup Instructions](#setup-instructions)
3. [API Endpoints](#api-endpoints)
4. [Environment Variables](#environment-variables)
5. [Running Migrations](#running-migrations)
6. [Testing the API](#testing-the-api)
7. [License](#license)

---
### Project Structure
```
user-microservice/
├── main.go
├── models/
│ └── user.go
├── handlers/
│ └── user_handler.go
├── database/
│ └── database.go
├── migrations/
│ └── migrations.go
└── .env
```

---

## Setup Instructions

### Prerequisites

- Go 1.20 or higher
- PostgreSQL
- Git

### Steps

1. **Clone the repository:**
   ```bash
   git clone https://github.com/allpyjs/go-user-microservice-assessment.git
   cd go-user-microservice-assessment
2. **Install dependences**
    ```bash
    go mod download
   ```
3. **Set up the database:**
4. **UPdate environment variables:**
    ```
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=password
    DB_NAME=userdb
    DB_PORT=5432
    ```