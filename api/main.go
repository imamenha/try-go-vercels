package handler

import (
	"log"
	"try-go-vercels/app/di"

	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// Handler is the exported function that Vercel will use
func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()

	handler := di.InitializeHandler()
	handler.RegisterRoutes(app)

	adaptor.FiberApp(app)(w, r)
}

func main() {
	app := fiber.New()

	handler := di.InitializeHandler()
	handler.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
