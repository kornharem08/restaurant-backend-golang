package services

import "gorm.io/gorm"

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NewUserRequest struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type UserService interface {
	GetAll() ([]UserResponse, error)
	Create(newUser NewUserRequest) (*UserResponse, error)
	Login(loginRequest LoginRequest) (*LoginResponse, error)
}
