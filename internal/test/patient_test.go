package test

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "makerble-assessment/internal/model"
    "makerble-assessment/internal/repository"
    "makerble-assessment/internal/service"
)

func setupTestDB(t *testing.T) *gorm.DB {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }
    db.AutoMigrate(&model.Patient{})
    return db
}

func TestPatientService_Create(t *testing.T) {
    db := setupTestDB(t)
    repo := repository.NewPatientRepository(db)
    svc := service.NewPatientService(repo)

    input := service.CreatePatientInput{
        FirstName:   "John",
        LastName:    "Doe",
        DateOfBirth: "1990-01-01T00:00:00Z",
        Gender:      "Male",
        Contact:     "1234567890",
        Address:     "123 Main St",
    }

    patient, err := svc.Create(input)
    assert.NoError(t, err)
    assert.Equal(t, input.FirstName, patient.FirstName)
    assert.Equal(t, input.LastName, patient.LastName)
    assert.Equal(t, input.Contact, patient.Contact)
    assert.Equal(t, input.Address, patient.Address)
}

func TestPatientService_Get(t *testing.T) {
    db := setupTestDB(t)
    repo := repository.NewPatientRepository(db)
    svc := service.NewPatientService(repo)

    patient := model.Patient{
        FirstName:   "Jane",
        LastName:    "Doe",
        DateOfBirth: time.Now(),
        Gender:      "Female",
        Contact:     "0987654321",
        Address:     "456 Elm St",
    }
    db.Create(&patient)

    result, err := svc.Get(patient.ID)
    assert.NoError(t, err)
    assert.Equal(t, patient.FirstName, result.FirstName)
    assert.Equal(t, patient.LastName, result.LastName)
}