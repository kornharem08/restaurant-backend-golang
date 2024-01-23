package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kornharem08/society-shop/databases"
	"github.com/kornharem08/society-shop/handlers"
	"github.com/kornharem08/society-shop/middlewares"
	"github.com/kornharem08/society-shop/repositories"
	"github.com/kornharem08/society-shop/services"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env")
	}
	db := databases.ConnectDB()
	secretKey := os.Getenv("SECRET_KEY")
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	app := fiber.New()
	jwt := middlewares.NewAuthMiddleware(secretKey)
	app.Get("/users", jwt, userHandler.GetAll)
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
