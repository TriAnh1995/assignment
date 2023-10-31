# **Project: friend-management**
### How to Run:
- Setup: make pg, make setup
- Run app: make run

### **Technology**:
- Using Go 1.19
- PostgreSQL
- DB migration
- sqlBoiler
- mockery

### Project structure
- Workflow: Request => Routes => Handler => Controller => Repository => Database

- Three layers model:
    + internal/handler: Get request from httpRequest, decode, validate, call controller, write httpResponse
    + internal/controller: Handle business logic, call repositories
    + internal/repository: Data access layer

### API ENDPOINTS(localhost:5432)
1. Create User:
   Path: http:/localhost:5432/users
   Method: Post.

    Example:
   {
        "name":"test",
        "email":"test@example.com"
   }