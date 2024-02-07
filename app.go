package main

import (
	"zoc-api/config"
	"zoc-api/database"
	"zoc-api/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadEnv()
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Use(logger.New())

	v1 := app.Group("/api/v1")

	customers := v1.Group("/customers")
	customers.Get("/", handlers.GetCustomers)
	customers.Get("/:id", handlers.GetCustomer)
	customers.Post("/", handlers.CreateCustomer)
	customers.Put("/:id", handlers.UpdateCustomer)
	customers.Delete("/:id", handlers.DeleteCustomer)

	app.Listen(":3000")
}