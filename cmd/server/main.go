package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "makerble-assessment/internal/config"
    "makerble-assessment/internal/handler"
    "makerble-assessment/internal/middleware"
    "makerble-assessment/internal/repository"
    "makerble-assessment/internal/service"
    _ "makerble-assessment/docs"
)

// @title Makerble Assessment API
// @version 1.0
// @description Hospital management API for Makerble internship
// @host localhost:8080
// @BasePath /
func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    db, err := config.InitDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    r := gin.Default()

    userRepo := repository.NewUserRepository(db)
    patientRepo := repository.NewPatientRepository(db)
    authService := service.NewAuthService(userRepo)
    patientService := service.NewPatientService(patientRepo)
    authHandler := handler.NewAuthHandler(authService)
    patientHandler := handler.NewPatientHandler(patientService)

    r.POST("/login", authHandler.Login)

    receptionist := r.Group("/api/receptionist").Use(middleware.AuthMiddleware(authService, "receptionist"))
    {
        receptionist.POST("/patients", patientHandler.Create)
        receptionist.GET("/patients", patientHandler.List)
        receptionist.GET("/patients/:id", patientHandler.Get)
        receptionist.PUT("/patients/:id", patientHandler.Update)
        receptionist.DELETE("/patients/:id", patientHandler.Delete)
    }

    doctor := r.Group("/api/doctor").Use(middleware.AuthMiddleware(authService, "doctor"))
    {
        doctor.GET("/patients", patientHandler.List)
        doctor.GET("/patients/:id", patientHandler.Get)
        doctor.PUT("/patients/:id", patientHandler.UpdateMedicalHistory)
    }

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Failed to run server:", err)
    }
}