package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/DavAnders/SkillTogether/bot/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Load bot token
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		fmt.Println("DISCORD_BOT_TOKEN is not set")
		return
	}

	// Create Discord session
	discordSession, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	discordSession.AddHandler(messageCreate)

	err = discordSession.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Discord bot is running, press CTRL+C to exit")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop // Block until CTRL+C is pressed

	discordSession.Close()
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return // Ignore bot's own messages
	}

	// Check if this message is part of an ongoing interaction
	if handled := commands.HandleOngoingInteraction(session, message); handled {
		return // If the message was part of an ongoing interaction, stop further processing
	}

	// Handle new commands
	if strings.HasPrefix(message.Content, "!") {
		commands.HandleCommand(session, message)
	}
}
