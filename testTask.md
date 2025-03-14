**Go Developer Interview Task**

**Duration:** 1 Hour

### **Objective**
The goal of this task is to assess the candidate's ability to set up a Go project with a microservices architecture, implement CRUD API endpoints, connect to a PostgreSQL database, and optionally integrate OAuth authentication.

---

### **Task Description**

You are required to build a simple microservices-based application using Go. The application should include a basic service for managing users.

#### **Requirements:**
1. **Project Setup:**
   - Initialize a Go module.
   - Use a standard project structure for a microservices-based application.
   - Implement a basic HTTP server using a Go framework of your choice (e.g., Gin, Echo, Fiber).

2. **CRUD API Endpoints for User Management:**
   - Create a `User` entity with the following fields:
     - `ID` (UUID, Primary Key)
     - `Name` (string)
     - `Email` (string, unique)
     - `CreatedAt` (timestamp)
   - Implement RESTful API endpoints:
     - `POST /users` – Create a new user
     - `GET /users/{id}` – Get user details by ID
     - `PUT /users/{id}` – Update user details
     - `DELETE /users/{id}` – Delete a user

3. **Database Connectivity with PostgreSQL:**
   - Use GORM or the standard `database/sql` package to connect to a PostgreSQL database.
   - Write migration scripts to create the `users` table.

4. **(Optional) OAuth Authentication:**
   - Secure the endpoints using OAuth 2.0.
   - Implement authentication using a library like `oauth2` or integrate with an identity provider (e.g., Google OAuth, Auth0).

---

### **Technical Guidelines**
- Use Go modules for dependency management.
- Organize the project following clean architecture principles.
- Implement proper error handling and validation.
- Use environment variables for configuration (e.g., database credentials).
- Ensure that the code is well-structured, readable, and maintainable.

---

### **Deliverables**
- A GitHub repository (or ZIP file) containing the source code.
- A `README.md` file with instructions on how to run the project.
- A Postman collection (optional) for testing the APIs.

---







### **Evaluation Criteria**
- **Code Quality & Best Practices** – Follows Go best practices and clean architecture principles.
- **Microservices Setup** – Properly structured project with clear separation of concerns.
- **API Implementation** – Functional and RESTful CRUD endpoints.
- **Database Integration** – Efficient use of PostgreSQL with migrations.
- **Bonus: OAuth Integration** – Secure API authentication and authorization.

---

### **Additional Notes**
- The task is expected to be completed within an hour.
- Focus on functionality, but also on writing clean and maintainable code.
- If OAuth integration cannot be completed, ensure that the basic structure is in place for future extension.

Good luck!
