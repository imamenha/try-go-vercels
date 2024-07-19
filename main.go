package main

import (
	"log"
	"try-go-vercels/app/di"

	"github.com/gofiber/fiber/v2"
)

// Function exported to Vercel
func Handler() *fiber.App {
	app := fiber.New()

	handler := di.InitializeHandler()
	handler.RegisterRoutes(app)

	return app
}

func main() {
	app := Handler()
	log.Fatal(app.Listen(":3000"))
}
