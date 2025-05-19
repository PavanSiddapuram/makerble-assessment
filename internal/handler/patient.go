package handler

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "makerble-assessment/internal/service"
)

type PatientHandler struct {
    service *service.PatientService
}

func NewPatientHandler(service *service.PatientService) *PatientHandler {
    return &PatientHandler{service: service}
}

// Create godoc
// @Security BearerAuth
// @Summary Create a patient
// @Description Create a new patient (receptionist only)
// @Tags receptionist
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param patient body service.CreatePatientInput true "Patient details"
// @Success 201 {object} service.PatientResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/receptionist/patients [post]
func (h *PatientHandler) Create(c *gin.Context) {
    var input service.CreatePatientInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    patient, err := h.service.Create(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, patient)
}

// List godoc
// @Security BearerAuth
// @Summary List patients
// @Description Get a list of patients
// @Tags receptionist,doctor
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} service.PatientResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/receptionist/patients [get]
// @Router /api/doctor/patients [get]
func (h *PatientHandler) List(c *gin.Context) {
    patients, err := h.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, patients)
}

// Get godoc
// @Security BearerAuth
// @Summary Get a patient
// @Description Get a patient by ID
// @Tags receptionist,doctor
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Patient ID"
// @Success 200 {object} service.PatientResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/receptionist/patients/{id} [get]
// @Router /api/doctor/patients/{id} [get]
func (h *PatientHandler) Get(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    patient, err := h.service.Get(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
        return
    }

    c.JSON(http.StatusOK, patient)
}

// Update godoc
// @Security BearerAuth
// @Summary Update a patient
// @Description Update a patient's details (receptionist only)
// @Tags receptionist
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Patient ID"
// @Param patient body service.UpdatePatientInput true "Patient details"
// @Success 200 {object} service.PatientResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/receptionist/patients/{id} [put]
func (h *PatientHandler) Update(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var input service.UpdatePatientInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    patient, err := h.service.Update(uint(id), input)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, patient)
}

// Delete godoc
// @Security BearerAuth
// @Summary Delete a patient
// @Description Delete a patient by ID (receptionist only)
// @Tags receptionist
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Patient ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/receptionist/patients/{id} [delete]
func (h *PatientHandler) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.service.Delete(uint(id)); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}

// UpdateMedicalHistory godoc
// @Security BearerAuth
// @Summary Update a patient's medical history
// @Description Update a patient's medical history (doctor only)
// @Tags doctor
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Patient ID"
// @Param medical_history body service.MedicalHistoryInput true "Medical history"
// @Success 200 {object} service.PatientResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/doctor/patients/{id} [put]
func (h *PatientHandler) UpdateMedicalHistory(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var input service.MedicalHistoryInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    patient, err := h.service.UpdateMedicalHistory(uint(id), input.MedicalHistory)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, patient)
}