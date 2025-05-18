package model

import (
    "time"
    "gorm.io/gorm"
)

type Patient struct {
    gorm.Model
    FirstName      string    `gorm:"not null"`
    LastName       string    `gorm:"not null"`
    DateOfBirth    time.Time `gorm:"not null"`
    Gender         string    `gorm:"not null"`
    Contact        string
    Address        string
    MedicalHistory string
}