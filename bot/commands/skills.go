package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
}

var (
	userStates  = make(map[string]string)
	tempStorage = make(map[string]string)
	statesMutex = &sync.Mutex{}
)

// PostRequest represents the JSON request body for adding a skill.
type PostRequest struct {
	DiscordID   string `json:"discord_id"`
	Description string `json:"skill_description"`
}

// InterestPostRequest represents the JSON request body for adding an interest.
type InterestPostRequest struct {
	DiscordID string `json:"discord_id"`
	Interest  string `json:"interest"`
}

func clearUserState(userID string) {
	statesMutex.Lock()
	delete(userStates, userID)
	delete(tempStorage, userID)
	statesMutex.Unlock()
}

// HandleOngoingInteraction checks if the user is in an ongoing interaction and handles the message accordingly.
func HandleOngoingInteraction(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	statesMutex.Lock()
	state, inConversation := userStates[m.Author.ID]
	statesMutex.Unlock()

	if !inConversation {
		return false // No ongoing interaction for this user
	}

	// Handle the message based on the current state
	switch state {
	case "expecting_skill":
		handleExpectingSkill(s, m)
	case "confirm_skill":
		handleConfirmSkill(s, m)
	case "expecting_interest":
		handleExpectingInterest(s, m)
	case "confirm_interest":
		handleConfirmInterest(s, m)
		// Add other states and functions if context grows
	}

	return true // Message was part of an ongoing interaction
}

// AddSkill starts the process of adding a new skill.
func AddSkill(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Setup state + initial prompt only
	statesMutex.Lock()
	// Check if already in a conversation to avoid resetting states
	if _, exists := userStates[message.Author.ID]; !exists {
		userStates[message.Author.ID] = "expecting_skill"
		session.ChannelMessageSend(message.ChannelID, "Please reply with the skill you want to add, or reply with 'cancel' to cancel.")
	}
	statesMutex.Unlock()
}

// AddInterest starts the process of adding a new interest.
func AddInterest(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Setup state + initial prompt only
	statesMutex.Lock()
	// Check if already in a conversation to avoid resetting states
	if _, exists := userStates[message.Author.ID]; !exists {
		userStates[message.Author.ID] = "expecting_interest"
		session.ChannelMessageSend(message.ChannelID, "Please reply with the interest you want to add, or reply with 'cancel' to cancel.")
	}
	statesMutex.Unlock()
}

// HandleSkillState handles the user's skill state.
// This function is called when the user is in the "expecting_skill" or "confirm_skill" state.
func HandleSkillState(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	statesMutex.Lock()
	state, ok := userStates[message.Author.ID]
	statesMutex.Unlock()

	if ok {
		switch state {
		case "expecting_skill":
			handleExpectingSkill(session, message)
		case "confirm_skill":
			handleConfirmSkill(session, message)
		}
	}
}

// HandleInterestState handles the user's interest state.
// This function is called when the user is in the "expecting_interest" or "confirm_interest" state.
func HandleInterestState(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	statesMutex.Lock()
	state, ok := userStates[message.Author.ID]
	statesMutex.Unlock()

	if ok {
		switch state {
		case "expecting_interest":
			handleExpectingInterest(session, message)
		case "confirm_interest":
			handleConfirmInterest(session, message)
		}
	}
}

func handleExpectingSkill(session *discordgo.Session, message *discordgo.MessageCreate) {
	content := strings.ToLower(strings.TrimSpace(message.Content))
	if content == "cancel" {
		session.ChannelMessageSend(message.ChannelID, "Skill addition cancelled.")
		clearUserState(message.Author.ID)
		return
	}
	session.ChannelMessageSend(message.ChannelID, "Would you like to post this as a skill available for trade? Reply 'yes' to confirm or 'no' to cancel.")
	statesMutex.Lock()
	userStates[message.Author.ID] = "confirm_skill"
	tempStorage[message.Author.ID] = message.Content // Skill description
	statesMutex.Unlock()
}

func handleConfirmSkill(session *discordgo.Session, message *discordgo.MessageCreate) {
	content := strings.ToLower(strings.TrimSpace(message.Content))
	if content == "yes" {
		description := tempStorage[message.Author.ID]

		data := PostRequest{
			DiscordID:   message.Author.ID,
			Description: description,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Error processing request: unable to marshal JSON.")
			return
		}

		baseURL := os.Getenv("API_URL")
		if baseURL == "" {
			session.ChannelMessageSend(message.ChannelID, "API URL is not set.")
			return
		}

		// Create an HTTP client and request
		client := &http.Client{}
		req, err := http.NewRequest("POST", baseURL+"/bot/skills", bytes.NewBuffer(jsonData))
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Failed to create request: "+err.Error())
			return
		}

		// Add the API key to the request header
		apiKey := os.Getenv("MY_API_KEY")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API-Key", apiKey)

		// Execute the POST request
		resp, err := client.Do(req)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Failed to post skill due to an error: %v", err))
			return
		}
		defer resp.Body.Close()

		// Check response status code
		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body) // Reading the body for additional error context
			session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Failed to post skill. Status: %s, Response: %s", resp.Status, string(bodyBytes)))
			return
		}

		session.ChannelMessageSend(message.ChannelID, "Skill posted successfully.")
		clearUserState(message.Author.ID)
	} else if content == "no" {
		session.ChannelMessageSend(message.ChannelID, "Skill addition cancelled.")
		clearUserState(message.Author.ID)
	} else {
		session.ChannelMessageSend(message.ChannelID, "Invalid response. Please reply 'yes' to confirm or 'no' to cancel.")
	}
}

func handleExpectingInterest(session *discordgo.Session, message *discordgo.MessageCreate) {
	content := strings.ToLower(strings.TrimSpace(message.Content))
	if content == "cancel" {
		session.ChannelMessageSend(message.ChannelID, "Interest addition cancelled.")
		clearUserState(message.Author.ID)
		return
	}
	session.ChannelMessageSend(message.ChannelID, "Would you like to post this as an interest? Reply 'yes' to confirm or 'no' to cancel.")
	statesMutex.Lock()
	userStates[message.Author.ID] = "confirm_interest"
	tempStorage[message.Author.ID] = message.Content // Interest description
	statesMutex.Unlock()
}

func handleConfirmInterest(session *discordgo.Session, message *discordgo.MessageCreate) {
	content := strings.ToLower(strings.TrimSpace(message.Content))
	if content == "yes" {
		interest := tempStorage[message.Author.ID]

		// Prepare the POST request to API
		data := InterestPostRequest{
			DiscordID: message.Author.ID,
			Interest:  interest,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Error processing request: unable to marshal JSON.")
			return
		}

		baseURL := os.Getenv("API_URL")
		if baseURL == "" {
			session.ChannelMessageSend(message.ChannelID, "API URL is not set.")
			return
		}

		// Create an HTTP client and request
		client := &http.Client{}
		req, err := http.NewRequest("POST", baseURL+"/bot/interests", bytes.NewBuffer(jsonData))
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Failed to create request: "+err.Error())
			return
		}

		// Add the API key to the request header
		apiKey := os.Getenv("MY_API_KEY")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API-Key", apiKey)

		// Execute the POST request
		resp, err := client.Do(req)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Failed to post interest due to an error: %v", err))
			return
		}
		defer resp.Body.Close()

		// Check response status code
		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body) // Reading the body for additional error context
			session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Failed to post interest. Status: %s, Response: %s", resp.Status, string(bodyBytes)))
			return
		}

		session.ChannelMessageSend(message.ChannelID, "Interest posted successfully.")
		clearUserState(message.Author.ID)
	} else if content == "no" {
		session.ChannelMessageSend(message.ChannelID, "Interest addition cancelled.")
		clearUserState(message.Author.ID)
	} else {
		session.ChannelMessageSend(message.ChannelID, "Invalid response. Please reply 'yes' to confirm or 'no' to cancel.")
	}
}

// ListSkills - TODO
func ListSkills(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "Listing skills")
	// need to fill out
}

// ListInterests - TODO
func ListInterests(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "Listing interests")
	// need to fill out
}

// HandleCommand handles the bot commands.
func HandleCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	frontendURL := os.Getenv("FRONTEND_URL")
	if strings.HasPrefix(message.Content, "!") {
		// Parse + direct
		command := strings.ToLower(strings.TrimSpace(message.Content))
		switch {
		case command == "!site":
			msg := fmt.Sprintf("Visit the site at %s to view the dashboard.", frontendURL)
			session.ChannelMessageSend(message.ChannelID, msg)
		case command == "!addskill":
			AddSkill(session, message)
		case command == "!addinterest":
			AddInterest(session, message)
		case command == "!ping":
			session.ChannelMessageSend(message.ChannelID, "Pong!")
		case command == "!help":
			session.ChannelMessageSend(message.ChannelID, "List of commands:\n`!addskill` - Start adding a new skill.\n`!addinterest` - Start adding a new interest.\n`!ping` - Responds with 'Pong!'.\n `!site` - Displays the link to the website dashboard. \n`!help` - Shows this message.")
		default:
			session.ChannelMessageSend(message.ChannelID, "Unknown command. Type !help for a list of commands.")
		}
	}
}
