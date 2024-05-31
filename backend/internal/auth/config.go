package auth

import (
	"log"
	"os"
	"strings"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/joho/godotenv"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
)


var DiscordOAuth2Config *oauth2.Config

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return strings.TrimSpace(value)
}

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    clientID := getEnv("CLIENT_ID", "")
    clientSecret := getEnv("CLIENT_SECRET", "")
    redirectURL := getEnv("REDIRECT_URL", "")

    if clientID == "" || clientSecret == "" || redirectURL == "" {
        log.Fatal("One or more environment variables are not set.")
    }

    DiscordOAuth2Config = &oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURL:  redirectURL,
        Scopes:       []string{"identify", "email"},
        Endpoint:     discord.Endpoint,
    }
}

type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Email		  string `json:"email"`
	Avatar     	  string `json:"avatar"`
}

// Construct the avatar URL
func (user *DiscordUser) AvatarURL() string {
	if user.Avatar == "" {
		return "" // Can return default avatar URL here later
	}
	return "https://cdn.discordapp.com/avatars/" + user.ID + "/" + user.Avatar + ".png"
}

type AuthHandler struct {
	Queries *db.Queries
}

func NewAuthHandler(queries *db.Queries) *AuthHandler {
	return &AuthHandler{
		Queries: queries,
	}
}