# 📚 Book Management API with Go Fiber and JWT

A simple RESTful API for managing book data using **Golang**, **Fiber**, **JWT Authentication**, and **MySQL**.  
This project is made for learning purposes, especially for SMK students who want to practice building a secure CRUD API.

---

## 📌 Features

✅ Register & Login to get a **JWT Token**  
✅ CRUD (Create, Read, Update, Delete) for books  
✅ JWT & Bearer Token protection for all book routes  
✅ MySQL database integration  
✅ Clean and simple project structure with clear explanations

---

## ⚙️ Tech Stack

- **Golang**
- **Fiber** (fast HTTP web framework)
- **GORM** (ORM for MySQL)
- **JWT** (JSON Web Token)

---

## 🗂️ Project Structure

```plaintext
.
├── config/          # Database connection
│   └── database.go
├── controllers/     # Logic for Register, Login, CRUD Books
│   ├── authController.go
│   └── bookController.go
├── middleware/      # JWT Middleware
│   └── jwtMiddleware.go
├── models/          # User and Book models
│   ├── user.go
│   └── book.go
├── routes/          # Route definitions
│   └── routes.go
├── main.go          # Main entry point
├── go.mod           # Go modules
└── go.sum
