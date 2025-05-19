package test

import (
    "fmt"
    "os"
    "testing"
    "time" // Add time import
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "makerble-assessment/internal/model"
    "makerble-assessment/internal/repository"
    "makerble-assessment/internal/service"
)

func TestPatientService_Create(t *testing.T) {
    if err := godotenv.Load(); err != nil {
        t.Fatal("Error loading .env file")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    if err := db.AutoMigrate(&model.Patient{}); err != nil {
        t.Fatalf("Failed to migrate test database: %v", err)
    }

    repo := repository.NewPatientRepository(db)
    svc := service.NewPatientService(repo)

    input := service.CreatePatientInput{
        FirstName:   "Jane",
        LastName:    "Doe",
        DateOfBirth: "1995-05-05T00:00:00Z", // Use string
        Gender:      "Female",
        Contact:     "9876543210",
        Address:     "456 Elm St",
    }

    patient, err := svc.Create(input) // Pass input, not &input
    if err != nil {
        t.Fatalf("Failed to create patient: %v", err)
    }

    if patient.ID == 0 {
        t.Error("Patient ID should not be zero after creation")
    }
}

func TestPatientService_Get(t *testing.T) {
    if err := godotenv.Load(); err != nil {
        t.Fatal("Error loading .env file")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    if err := db.AutoMigrate(&model.Patient{}); err != nil {
        t.Fatalf("Failed to migrate test database: %v", err)
    }

    repo := repository.NewPatientRepository(db)
    svc := service.NewPatientService(repo)

    dob, _ := time.Parse(time.RFC3339, "1995-05-05T00:00:00Z")
    patient := model.Patient{
        FirstName:   "Jane",
        LastName:    "Doe",
        DateOfBirth: dob,
        Gender:      "Female",
        Contact:     "9876543210",
        Address:     "456 Elm St",
    }

    if err := db.Create(&patient).Error; err != nil {
        t.Fatalf("Failed to seed test patient: %v", err)
    }

    retrieved, err := svc.Get(patient.ID)
    if err != nil {
        t.Fatalf("Failed to get patient: %v", err)
    }

    if retrieved.ID != patient.ID || retrieved.FirstName != patient.FirstName {
        t.Errorf("Retrieved patient does not match: got %+v, want %+v", retrieved, patient)
    }
}