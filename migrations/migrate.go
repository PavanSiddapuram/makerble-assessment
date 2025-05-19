package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "makerble-assessment/internal/model"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    if err := db.AutoMigrate(&model.User{}, &model.Patient{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
    users := []model.User{
        {Email: "recep@example.com", Password: string(password), Role: "receptionist"},
        {Email: "doc@example.com", Password: string(password), Role: "doctor"},
    }

    for _, user := range users {
        if err := db.FirstOrCreate(&user, model.User{Email: user.Email}).Error; err != nil {
            log.Printf("Failed to create user %s: %v", user.Email, err)
        }
    }
}