Makerble Assessment
Hospital management API for Makerble Golang internship.
Setup

Install Go and MySQL.
Create database: CREATE DATABASE makerble_db;
Update .env with MySQL credentials.
Run migrations: go run migrations/migrate.go
Start server: go run ./cmd/server
Access Swagger: http://localhost:8080/swagger/index.html

Endpoints

POST /login: Authenticate user
Receptionist APIs: /api/receptionist/patients (CRUD)
Doctor APIs: /api/doctor/patients (view/update medical history)

Running Tests
go test ./internal/test

Submission

GitHub: [your-repo-url]
Video demo: [link]
Documentation: [Google Doc link]

