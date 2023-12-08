package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:3000/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

func RedirectGoogle(c *fiber.Ctx) error {
	url := googleOauthConfig.AuthCodeURL("state")
	return c.Redirect(url)
}

func CallbackGoogle(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return c.Status(400).SendString("Failed to exchange token")
	}
	fmt.Println(token)
	return c.SendString("Success")
}


