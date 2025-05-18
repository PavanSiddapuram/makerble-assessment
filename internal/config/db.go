package config

import (
    "os"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "makerble-assessment/internal/model"
)

func InitDB() (*gorm.DB, error) {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, password, host, port, dbname)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    db.AutoMigrate(&model.User{}, &model.Patient{})
    return db, nil
}