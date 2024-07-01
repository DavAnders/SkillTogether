package auth

import (
	"log"
	"os"
	"strings"

	"github.com/DavAnders/SkillTogether/backend/db"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"

	"github.com/joho/godotenv"
)

// DiscordOAuth2Config is a global variable to store OAuth2 configuration for Discord.
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

// DiscordUser represents a user's Discord information.
type DiscordUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

// AvatarURL returns the URL of the user's avatar.
// If the user has no avatar, it returns an empty string.
func (user *DiscordUser) AvatarURL() string {
	if user.Avatar == "" {
		return "" // Can return default avatar URL here later
	}
	return "https://cdn.discordapp.com/avatars/" + user.ID + "/" + user.Avatar + ".png"
}

// AuthHandler handles authentication-related queries.
type AuthHandler struct {
	Queries *db.Queries
}

// NewAuthHandler initializes a new AuthHandler with the provided queries.
func NewAuthHandler(queries *db.Queries) *AuthHandler {
	return &AuthHandler{
		Queries: queries,
	}
}
