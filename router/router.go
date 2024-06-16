package router

import (
	"api-fiber-gorm/handler"
	"api-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.CreateUser)

	// User
	user := api.Group("/user")
	user.Get("/all", middleware.Protected(), handler.GetUsers)
	user.Get("/:id", middleware.Protected(), handler.GetUser)
	user.Put("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Note
	note := api.Group("/note")
	note.Get("/all", middleware.Protected(), handler.GetAllNotes)
	note.Get("/:id", middleware.Protected(), handler.GetNote)
	note.Post("/create", middleware.Protected(), handler.CreateNote)
	note.Put("/:id", middleware.Protected(), handler.UpdateNote)
	note.Delete("/:id", middleware.Protected(), handler.DeleteNote)
}
