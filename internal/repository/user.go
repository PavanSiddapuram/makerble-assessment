package repository

import (
    "gorm.io/gorm"
    "makerble-assessment/internal/model"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (model.User, error) {
    var user model.User
    err := r.db.Where("email = ?", email).First(&user).Error
    return user, err
}