package fiber

import (
	"github.com/gofiber/fiber/v2"
	"taskManagement/middleware"
)

// New create a new fiber route for the apis
func New() *fiber.App {
	app := fiber.New()
	app.Use(middleware.Recover) // convert panics to 5XX responses -- handle any panic that we get from server response
	return app
}
