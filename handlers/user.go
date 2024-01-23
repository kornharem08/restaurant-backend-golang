package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kornharem08/society-shop/services"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	user, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   user,
	})

}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	request := &services.NewUserRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if request.Email == "" || request.Password == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Email and password are required"})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Name are required"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	request.Password = string(hashedPassword)

	response, err := h.service.Create(*request)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": fiber.StatusCreated,
		"data":   response,
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	request := &services.LoginRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if request.Email == "" || request.Password == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Email and password are required"})
	}

	response, err := h.service.Login(*request)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":       fiber.StatusOK,
		"access_token": response.Token,
	})

}
