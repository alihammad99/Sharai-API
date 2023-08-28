package middlewares

import (
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/contrib/jwt"
)

func JwtMiddleware() fiber.Handler {
    return jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte("secret")},
    })
}
