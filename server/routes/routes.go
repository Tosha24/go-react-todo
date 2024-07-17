package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tosha24/todo/controller"
	"github.com/tosha24/todo/middleware"
)

func MyRoutes(app *fiber.App) {
	// Auth routes
	app.Post("/api/login", controller.Login)
	app.Post("/api/register", controller.CreateUser)

	app.Post("/api", middleware.JWTMiddleware)

	// todo routes	
	app.Post("/api/todo", controller.AddTodo)
	app.Get("/api/todos", controller.GetAllTodos)
	app.Put("/api/todo/mark/:id", controller.MarkAsCompleted)
	app.Put("/api/todo/:id", controller.UpdateTodo)
	app.Delete("/api/todo/:id", controller.DeleteTodo)
}