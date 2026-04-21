# 🚀 Go Clean Architecture REST API

A production-ready REST API built with **Golang, Gin, PostgreSQL**, following **Clean Architecture principles**.

---

## 🧠 Architecture Overview

This project follows **Clean Architecture**, ensuring separation of concerns:

```
Router → Controller → Usecase → Repository → Database
                ↑
             Domain
```

- **Domain** → Core business models & interfaces
- **Usecase** → Business logic
- **Repository** → Database interaction
- **Controller** → HTTP layer (Gin)
- **Bootstrap** → App initialization

---

## ⚙️ Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT (Access + Refresh Tokens)
- **Containerization**: Docker

---

## 📁 Project Structure

```
.
├── cmd/                  # Entry point
├── api/
│   ├── controller/       # HTTP handlers
│   ├── middleware/       # JWT middleware
│   └── route/            # Route definitions
├── domain/               # Entities & interfaces
├── usecase/              # Business logic
├── repository/           # DB operations
├── bootstrap/            # App & DB setup
├── internal/             # Utilities (JWT, etc.)
├── docker-compose.yaml
├── Dockerfile
└── .env
```

---

## 🔐 Authentication Flow

```
Login → Access Token (15 min) + Refresh Token (7 days)
Access expires → Use Refresh Token → New Access Token
```

---

## 🛠️ Setup Instructions

### 1. Clone the repository

```bash
git clone <your-repo-url>
cd go-clean-api
```

---

### 2. Create `.env` file

```env
PORT=8080

DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=clean_api
DB_PORT=5432

JWT_SECRET=super-secret-key
```

---

### 3. Run with Docker (Recommended)

```bash
docker-compose up --build
```

---

### 4. Run locally (without Docker)

```bash
go mod tidy
go run cmd/main.go
```

---

## 🧪 API Endpoints

### 🔑 Auth

#### Signup

```http
POST /signup
```

Body:

```json
{
  "email": "user@gmail.com",
  "password": "123456"
}
```

---

#### Login

```http
POST /login
```

Response:

```json
{
  "access_token": "jwt",
  "refresh_token": "jwt"
}
```

---

#### Refresh Token

```http
POST /refresh
```

Body:

```json
{
  "refresh_token": "jwt"
}
```

---

### 📌 Tasks (Protected)

#### Create Task

```http
POST /tasks
Authorization: Bearer <access_token>
```

Body:

```json
{
  "title": "Learn Go Clean Architecture"
}
```

---

#### Get Tasks

```http
GET /tasks
Authorization: Bearer <access_token>
```

---

## 🔒 Middleware

- JWT Authentication Middleware
- Extracts `user_id` from token
- Protects private routes

---

## 🧠 Key Features

- Clean Architecture (layered design)
- Dependency Injection
- JWT Authentication (Access + Refresh)
- PostgreSQL with GORM
- Dockerized setup
- Scalable and testable codebase

---

## ⚠️ Future Improvements

- Refresh token rotation
- Logout endpoint (invalidate tokens)
- Role-based access control (RBAC)
- Rate limiting
- Redis caching
- CI/CD pipeline

---

## 👨‍💻 Author

**Your Name**

---

## 📄 License

This project is licensed under the MIT License.
