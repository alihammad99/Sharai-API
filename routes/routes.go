package routes

import (
	"github.com/gofiber/fiber/v2"
	"sharai-api/controllers/auth"
	"sharai-api/middlewares"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/auth/google", auth.RedirectGoogle)
	app.Get("/auth/google/callback", auth.CallbackGoogle)

	app.Get("/auth/facebook", auth.RedirectFacebook)
	app.Get("/auth/facebook/callback", auth.CallbackFacebook)
	app.Post("/login", auth.Login)
	app.Post("/sendotp", auth.SMSAuth)

    app.Use(middlewares.JwtMiddleware())
	app.Post("/restricted", auth.Restricted)
}
