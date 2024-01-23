package services

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/kornharem08/society-shop/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) GetAll() ([]UserResponse, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	userRes := []UserResponse{}
	for _, customer := range users {
		user := UserResponse{
			Name:  customer.Name,
			Email: customer.Email,
		}
		userRes = append(userRes, user)
	}

	return userRes, nil
}

func (s *userService) Create(request NewUserRequest) (*UserResponse, error) {
	user := repositories.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil
}

func (s *userService) Login(request LoginRequest) (*LoginResponse, error) {
	user, err := s.repo.FindByCredentials(request.Email, request.Password)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized)
	}

	token, err := CreateToken(user.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized)
	}

	return &LoginResponse{
		Token: token.AccessToken,
	}, nil
}

func CreateToken(email string) (Token, error) {
	var msgToken Token
	secretKey := os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return msgToken, err
	}
	msgToken.AccessToken = t
	return msgToken, nil
}
