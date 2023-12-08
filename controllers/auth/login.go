func loginHandler(c *fiber.Ctx) error {
    // Parse the JSON request body
    type LoginRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    var loginReq LoginRequest
    if err := c.BodyParser(&loginReq); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
    }

    // Check if the user exists in the database
    filter := bson.M{"email": loginReq.Email, "password": loginReq.Password}
    var user struct{}
    if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password"})
    }

    // Successful login
    return c.JSON(fiber.Map{"message": "Login successful"})
}
