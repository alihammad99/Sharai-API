package main

import (
    "github.com/gofiber/fiber/v2"
    "sharai-api/controllers/database"
    "sharai-api/routes"
    "log"
    "github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
    db.ConnectDB()
    app := fiber.New()
    routes.SetupRoutes(app)

    app.Listen(":5000")
}
