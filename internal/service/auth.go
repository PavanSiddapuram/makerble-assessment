package service

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "makerble-assessment/internal/repository"
)

type AuthService struct {
    userRepo  *repository.UserRepository
    jwtSecret string
}

type LoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type UserResponse struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
    Role  string `json:"role"`
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
    return &AuthService{
        userRepo:  userRepo,
        jwtSecret: "mysecret123456", // Should be in .env
    }
}

func (s *AuthService) Login(input LoginInput) (string, UserResponse, error) {
    user, err := s.userRepo.FindByEmail(input.Email)
    if err != nil {
        return "", UserResponse{}, errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        return "", UserResponse{}, errors.New("invalid credentials")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "role":    user.Role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(s.jwtSecret))
    if err != nil {
        return "", UserResponse{}, err
    }

    return tokenString, UserResponse{
        ID:    user.ID,
        Email: user.Email,
        Role:  user.Role,
    }, nil
}

func (s *AuthService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(s.jwtSecret), nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}