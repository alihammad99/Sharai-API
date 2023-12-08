package auth

import (
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
    var data map[string]string

    if err := c.BodyParser(&data); err != nil {
        return err
    }

    user := data["user"]
    pass := data["pass"]

    if user != "john" || pass != "doe" {
        return c.SendStatus(fiber.StatusUnauthorized)
    }

    claims := jwt.MapClaims{
        "name":  "John Doe",
        "admin": true,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    t, err := token.SignedString([]byte("secret"))
    if err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(fiber.Map{"token": t})
}

func Accessible(c *fiber.Ctx) error {
    return c.SendString("Accessible")
}

func Restricted(c *fiber.Ctx) error {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    name := claims["name"].(string)
    return c.SendString("Welcome " + name)
}
