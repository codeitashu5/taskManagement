package task

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"taskManagement/jwt"
	"taskManagement/middleware"
	"taskManagement/models"
	"taskManagement/validator"
	"time"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Serve(router fiber.Router) {
	router.Post("/register", h.registerUser) // register the user and return access token and refresh token
	router.Post("/login", h.login)

	router.Use(middleware.ParseToken)
	router.Post("/logout", h.logout)
	router.Route("/tasks/", func(taskRouter fiber.Router) {
		taskRouter.Post("/", h.createTask)          // Create a new task
		taskRouter.Delete("/:taskId", h.deleteTask) // Delete a task (with archived_at check)
		taskRouter.Get("/:taskId", h.getTask)       // Get a single task by ID
		taskRouter.Get("/", h.getAllTasks)          // Get all tasks
		taskRouter.Put("/:taskId", h.updateTask)    // Update a task
	})

}

// create a new user
func (h *Handler) registerUser(c *fiber.Ctx) error {
	var body RegisterBody
	if err := validator.Parse(c, nil, nil, &body); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userId, err := h.service.registerUser(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return jwt.SendJWT(c, models.User{
		ID: userId,
	})
}

// login using the email and password
func (h *Handler) login(c *fiber.Ctx) error {
	var body LoginRequest
	if err := validator.Parse(c, nil, nil, &body); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	user, err := h.service.login(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return jwt.SendJWT(c, user)
}

// crate task for the given user
func (h *Handler) createTask(c *fiber.Ctx) error {
	var body Request
	if err := validator.Parse(c, nil, nil, &body); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	claims := middleware.ClaimsFromContext(c)
	err := h.service.createTask(body, claims.UserID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(models.MessageResponse{
		Message: "Task created",
	})
}

// get task by task-id
func (h *Handler) getTask(c *fiber.Ctx) error {
	taskIDParam := c.Params("taskId")
	taskID, err := primitive.ObjectIDFromHex(taskIDParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid task ID")
	}

	task, err := h.service.getTask(taskID)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(task)
}

// get all task
func (h *Handler) getAllTasks(c *fiber.Ctx) error {
	claims := middleware.ClaimsFromContext(c)
	tasks, err := h.service.getAllTasks(claims.UserID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(tasks)
}

func (h *Handler) deleteTask(c *fiber.Ctx) error {
	taskIDParam := c.Params("taskId")
	taskID, err := primitive.ObjectIDFromHex(taskIDParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid task ID")
	}

	err = h.service.deleteTask(taskID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(models.MessageResponse{
		Message: "Task deleted",
	})
}

func (h *Handler) updateTask(c *fiber.Ctx) error {
	taskIDParam := c.Params("taskId")
	taskID, err := primitive.ObjectIDFromHex(taskIDParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid task ID")
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = h.service.updateTask(taskID, body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(models.MessageResponse{
		Message: "Task updated",
	})
}

func (h *Handler) logout(c *fiber.Ctx) error {
	// Clear the JWT cookie by setting an empty value and an expiration date in the past
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour), // Set expiration time to the past
	})

	return c.Status(http.StatusOK).SendString("Successfully logged out")
}
