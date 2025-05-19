Makerble Golang Assessment - Hospital API
Overview
This project is a RESTful API for a hospital management system, built with Go, Gin, GORM, and MySQL. It supports two user roles:

Receptionist: Create, list, get, update, and delete patient records.
Doctor: List, get, and update patient medical history.

The API uses JWT for authentication, Swagger for documentation, and includes unit tests for patient operations.
Prerequisites

Go 1.24.3
MySQL 8.0
Git
Loom (for demo video, optional)

Setup

Clone the Repository:
git clone https://github.com/PavanSiddapuram/makerble-assessment.git
cd makerble-assessment


Configure Environment:

Copy .env.example to .env:cp .env.example .env


Edit .env:DB_USER=root
DB_PASSWORD=<your_mysql_password>
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=makerble_db




Install Dependencies:
go mod download


Set Up Database:

Start MySQL and create makerble_db:CREATE DATABASE makerble_db;


Run server to apply migrations:go run ./cmd/server




Run the Server:
go run ./cmd/server


Access Swagger UI: http://localhost:8080/swagger/index.html



API Usage
Authentication

POST /login:curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"email":"recep@example.com","password":"password123"}'


Returns JWT token. Use recep@example.com (receptionist) or doc@example.com (doctor).



Receptionist Endpoints (Role: receptionist)

POST /api/receptionist/patients: Create patient.curl -X POST http://localhost:8080/api/receptionist/patients -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"first_name":"John","last_name":"Doe","date_of_birth":"1990-01-01T00:00:00Z","gender":"Male","contact":"1234567890","address":"123 Main St"}'


GET /api/receptionist/patients: List patients.
GET /api/receptionist/patients/: Get patient.
PUT /api/receptionist/patients/: Update patient.
DELETE /api/receptionist/patients/: Delete patient.

Doctor Endpoints (Role: doctor)

GET /api/doctor/patients: List patients.
GET /api/doctor/patients/: Get patient.
PUT /api/doctor/patients/: Update medical history.curl -X PUT http://localhost:8080/api/doctor/patients/<id> -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"medical_history":"Patient has asthma"}'



Swagger Notes

Authorize with <token> (without Bearer) in Swagger UI due to middleware workaround.
If issues, use curl or Postman with Bearer <token>.

Testing

Run unit tests:go test ./internal/test


Tests cover patient creation and retrieval. Ensure .env is configured.

Troubleshooting

Tests fail with “.env not found”:
Verify .env exists in project root.
Check .env format (no spaces around =).


Swagger 401 Unauthorized:
Use <token> in Swagger’s Authorize dialog.
Alternatively, test with curl using Bearer <token>.



Submission

GitHub: https://github.com/PavanSiddapuram/makerble-assessment
Demo Video: https://www.loom.com/share/d8a3de28854d45d18003edb6dc829efc?sid=6417c974-0071-42c8-8c35-ab38fbd0deac (shows login, patient creation, doctor view, tests)



Contact
For issues, contact Pavan Siddapuram via siddapurampavan9381@gmail.com
