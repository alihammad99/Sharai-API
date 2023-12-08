package auth

import (
	"fmt"
    "net/http"
    "encoding/json"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/facebook"
	"crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)

var facebookOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:5000/auth/facebook/callback",
	ClientID:     "1617159938770979",
	ClientSecret: "bbdd85df55bd086b9bb7200fc40fdb28",
	Scopes:       []string{"public_profile", "email", "user_gender", "user_birthday"},
	Endpoint:     facebook.Endpoint,
}



func RedirectFacebook(c *fiber.Ctx) error {
	url := facebookOauthConfig.AuthCodeURL("state")
	return c.Redirect(url)
}

func generateAppSecretProof(appSecret string, accessToken string) string {
    key := []byte(appSecret)
    msg := []byte(accessToken)

    mac := hmac.New(sha256.New, key)
    mac.Write(msg)
    return hex.EncodeToString(mac.Sum(nil))
}

func CallbackFacebook(c *fiber.Ctx) error {
    code := c.Query("code")
    token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
    if err != nil {
        return c.Status(400).SendString("Failed to exchange token")
    }
    
    // Get the access token
    accessToken := token.AccessToken

	 // Generate the appsecret_proof
	 appSecret := "bbdd85df55bd086b9bb7200fc40fdb28"
	 appSecretProof := generateAppSecretProof(appSecret, accessToken)
    
    // Make a request to the Facebook Graph API to get the user's profile information
	resp, err := http.Get("https://graph.facebook.com/me?fields=id,name,email,birthday,gender,picture&access_token=" + accessToken + "&appsecret_proof=" + appSecretProof)

    if err != nil {
        return c.Status(500).SendString("Failed to get user data")
    }
    defer resp.Body.Close()
    
    // Parse the response
    var userData map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
        return c.Status(500).SendString("Failed to parse user data")
    }
    
    fmt.Println(userData)
    return c.SendString("Success")
}

