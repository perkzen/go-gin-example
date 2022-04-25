# Go rest API

![image](https://user-images.githubusercontent.com/73199603/164996068-d5aefac6-db02-46e3-9203-e2a490727f3e.png)

Go rest api using the framework Gin, MongoDb database and JWT authentication.

Server is running at <a> http://localhost:8080/api/v1 </a>

### Authentication routes

```
POST /auth/register
POST /auth/login
```

### User routes

```
GET /users (get all users)
GET /users/profile (get current user info)
```

### Todo routes

```
POST /todo (create todo)
GET /todo (get all current users todo's)
GET /todo/:id (get todo by id)
PUT /toggle/:id (toggle todo item bh id)
```