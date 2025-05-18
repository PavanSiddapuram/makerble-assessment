package service

import (
    "errors"
    "time"
    "makerble-assessment/internal/model"
    "makerble-assessment/internal/repository"
)

type PatientService struct {
    repo *repository.PatientRepository
}

type CreatePatientInput struct {
    FirstName      string `json:"first_name" binding:"required"`
    LastName       string `json:"last_name" binding:"required"`
    DateOfBirth    string `json:"date_of_birth" binding:"required"`
    Gender         string `json:"gender" binding:"required,oneof=Male Female Other"`
    Contact        string `json:"contact"`
    Address        string `json:"address"`
}

type UpdatePatientInput struct {
    FirstName      string `json:"first_name"`
    LastName       string `json:"last_name"`
    DateOfBirth    string `json:"date_of_birth"`
    Gender         string `json:"gender" binding:"omitempty,oneof=Male Female Other"`
    Contact        string `json:"contact"`
    Address        string `json:"address"`
}

type MedicalHistoryInput struct {
    MedicalHistory string `json:"medical_history" binding:"required"`
}

type PatientResponse struct {
    ID             uint   `json:"id"`
    FirstName      string `json:"first_name"`
    LastName       string `json:"last_name"`
    DateOfBirth    string `json:"date_of_birth"`
    Gender         string `json:"gender"`
    Contact        string `json:"contact"`
    Address        string `json:"address"`
    MedicalHistory string `json:"medical_history"`
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
    return &PatientService{repo: repo}
}

func (s *PatientService) Create(input CreatePatientInput) (PatientResponse, error) {
    dob, err := time.Parse(time.RFC3339, input.DateOfBirth)
    if err != nil {
        return PatientResponse{}, errors.New("invalid date of birth")
    }

    patient := model.Patient{
        FirstName:   input.FirstName,
        LastName:    input.LastName,
        DateOfBirth: dob,
        Gender:      input.Gender,
        Contact:     input.Contact,
        Address:     input.Address,
    }

    if err := s.repo.Create(&patient); err != nil {
        return PatientResponse{}, err
    }

    return PatientResponse{
        ID:             patient.ID,
        FirstName:      patient.FirstName,
        LastName:       patient.LastName,
        DateOfBirth:    patient.DateOfBirth.Format(time.RFC3339),
        Gender:         patient.Gender,
        Contact:        patient.Contact,
        Address:        patient.Address,
        MedicalHistory: patient.MedicalHistory,
    }, nil
}

func (s *PatientService) List() ([]PatientResponse, error) {
    patients, err := s.repo.FindAll()
    if err != nil {
        return nil, err
    }

    var response []PatientResponse
    for _, patient := range patients {
        response = append(response, PatientResponse{
            ID:             patient.ID,
            FirstName:      patient.FirstName,
            LastName:       patient.LastName,
            DateOfBirth:    patient.DateOfBirth.Format(time.RFC3339),
            Gender:         patient.Gender,
            Contact:        patient.Contact,
            Address:        patient.Address,
            MedicalHistory: patient.MedicalHistory,
        })
    }

    return response, nil
}

func (s *PatientService) Get(id uint) (PatientResponse, error) {
    patient, err := s.repo.FindByID(id)
    if err != nil {
        return PatientResponse{}, err
    }

    return PatientResponse{
        ID:             patient.ID,
        FirstName:      patient.FirstName,
        LastName:       patient.LastName,
        DateOfBirth:    patient.DateOfBirth.Format(time.RFC3339),
        Gender:         patient.Gender,
        Contact:        patient.Contact,
        Address:        patient.Address,
        MedicalHistory: patient.MedicalHistory,
    }, nil
}

func (s *PatientService) Update(id uint, input UpdatePatientInput) (PatientResponse, error) {
    patient, err := s.repo.FindByID(id)
    if err != nil {
        return PatientResponse{}, err
    }

    if input.FirstName != "" {
        patient.FirstName = input.FirstName
    }
    if input.LastName != "" {
        patient.LastName = input.LastName
    }
    if input.DateOfBirth != "" {
        dob, err := time.Parse(time.RFC3339, input.DateOfBirth)
        if err != nil {
            return PatientResponse{}, errors.New("invalid date of birth")
        }
        patient.DateOfBirth = dob
    }
    if input.Gender != "" {
        patient.Gender = input.Gender
    }
    if input.Contact != "" {
        patient.Contact = input.Contact
    }
    if input.Address != "" {
        patient.Address = input.Address
    }

    if err := s.repo.Update(&patient); err != nil {
        return PatientResponse{}, err
    }

    return PatientResponse{
        ID:             patient.ID,
        FirstName:      patient.FirstName,
        LastName:       patient.LastName,
        DateOfBirth:    patient.DateOfBirth.Format(time.RFC3339),
        Gender:         patient.Gender,
        Contact:        patient.Contact,
        Address:        patient.Address,
        MedicalHistory: patient.MedicalHistory,
    }, nil
}

func (s *PatientService) Delete(id uint) error {
    return s.repo.Delete(id)
}

func (s *PatientService) UpdateMedicalHistory(id uint, medicalHistory string) (PatientResponse, error) {
    patient, err := s.repo.FindByID(id)
    if err != nil {
        return PatientResponse{}, err
    }

    patient.MedicalHistory = medicalHistory
    if err := s.repo.Update(&patient); err != nil {
        return PatientResponse{}, err
    }

    return PatientResponse{
        ID:             patient.ID,
        FirstName:      patient.FirstName,
        LastName:       patient.LastName,
        DateOfBirth:    patient.DateOfBirth.Format(time.RFC3339),
        Gender:         patient.Gender,
        Contact:        patient.Contact,
        Address:        patient.Address,
        MedicalHistory: patient.MedicalHistory,
    }, nil
}