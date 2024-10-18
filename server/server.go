package server

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"taskManagement/mongoClient"
	apiFiber "taskManagement/server/fiber"
	"taskManagement/task"
)

type Server struct {
	app         *fiber.App // Changed from fiber.Router to fiber.App
	userHandler *task.Handler
}

func SetUpRoutes() *Server {
	// Initialize Fiber app
	app := apiFiber.New()

	// Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected To Task Management")
	})

	// Get the database from the client (assuming mongoClient.MongoDB is set up correctly)
	db := mongoClient.MongoDB.Database("taskManagmentDb")

	// Set up user service and handler
	userService := task.NewService(task.NewRepository(db))

	userHandler := task.NewHandler(userService)
	// Attach the user handler routes
	app.Route("/users", userHandler.Serve)

	// Return the initialized server with the Fiber app
	return &Server{
		app:         app,
		userHandler: userHandler,
	}
}

func (svc *Server) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Fiber handles the HTTP server internally, so you don't need to manually set up an http.Server.
	return svc.app.Listen(":" + port)
}
