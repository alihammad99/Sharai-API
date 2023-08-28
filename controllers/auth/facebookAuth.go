package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"os"
)

var facebookOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:3000/auth/facebook/callback",
	ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
	ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
	Scopes:       []string{"public_profile"},
	Endpoint:     facebook.Endpoint,
}



func RedirectFacebook(c *fiber.Ctx) error {
	url := facebookOauthConfig.AuthCodeURL("state")
	return c.Redirect(url)
}

func CallbackFacebook(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return c.Status(400).SendString("Failed to exchange token")
	}
	fmt.Println(token)
	return c.SendString("Success")
}
