package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kornharem08/society-shop/databases"
	"github.com/kornharem08/society-shop/handlers"
	"github.com/kornharem08/society-shop/repositories"
	"github.com/kornharem08/society-shop/services"
)

func main() {
	db := databases.ConnectDB()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	app := fiber.New()
	app.Get("/users", userHandler.GetAll)
	app.Post("/users", userHandler.Create)
	app.Post("/login", userHandler.Login)
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": fiber.StatusOK,
			"data":   "Hello World",
		})
	})
	app.Listen(":8080")
}
