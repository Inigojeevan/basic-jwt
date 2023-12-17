package main

import (
	"log"

	"github.com/Inigojeevan/jwt/config"
	"github.com/Inigojeevan/jwt/handlers"
	"github.com/Inigojeevan/jwt/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/login", handlers.Login)

	jwt := middleware.Auth(config.Secret)

	app.Get("/protected", jwt, handlers.ProtectedRoute)

	log.Fatal(app.Listen(":3000"))
}
