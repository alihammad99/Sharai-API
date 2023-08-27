package main

import (
    "github.com/gofiber/fiber/v2"
    "sharai-api/controllers"
    "sharai-api/middlewares"
)

func main() {
    app := fiber.New()

    // Login route
    app.Post("/login", controllers.Login)

    // Unauthenticated route
    app.Get("/", controllers.Accessible)

    // JWT Middleware
    app.Use(middlewares.JwtMiddleware())

    // Restricted Routes
    app.Post("/restricted", controllers.Restricted)

    app.Listen(":5000")
}
