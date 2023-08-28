package auth

import (
	"github.com/gofiber/fiber/v2"
	"sharai-api/middlewares"
)

const apiURL = "http://rest.d7networks.com/secure/send"

func SMSAuth(c *fiber.Ctx) error {
	var request struct {
		To string `json:"to"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	to := request.To
	otp := middlewares.GenerateOTP(4)

	err := middlewares.SendSMS(to, otp)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"status": "OTP sent successfully"})
}
