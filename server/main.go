package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tosha24/todo/config"
	"github.com/tosha24/todo/routes"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin, Authorization ",
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,

		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.MyRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
