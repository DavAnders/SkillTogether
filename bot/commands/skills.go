package commands

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	userStates = make(map[string]string)
	tempSkillStorage = make(map[string]string)
	statesMutex = &sync.Mutex{}
)

func clearUserState(userID string) {
    statesMutex.Lock()
    delete(userStates, userID)
    delete(tempSkillStorage, userID)
    statesMutex.Unlock()
}

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
    // Add other states and functions if context grows
    }

    return true // Message was part of an ongoing interaction
}

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
    tempSkillStorage[message.Author.ID] = message.Content  // Skill description
    statesMutex.Unlock()
}

func handleConfirmSkill(session *discordgo.Session, message *discordgo.MessageCreate) {
    content := strings.ToLower(strings.TrimSpace(message.Content))
    if content == "yes" {
        skillDescription := tempSkillStorage[message.Author.ID]
        fmt.Println("Skill posted: ", skillDescription) // Need to replace with actual posting logic
        session.ChannelMessageSend(message.ChannelID, "Skill posted.")
        clearUserState(message.Author.ID)
    } else if content == "no" {
        session.ChannelMessageSend(message.ChannelID, "Skill addition cancelled.")
        clearUserState(message.Author.ID)
    } else {
        session.ChannelMessageSend(message.ChannelID, "Invalid response. Please reply 'yes' to confirm or 'no' to cancel.")
    }
}

func ListSkills(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "Listing skills")
	// need to fill out
}

func HandleCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
    if strings.HasPrefix(message.Content, "!") {
        // Parse + direct
        command := strings.ToLower(strings.TrimSpace(message.Content))
        switch {
        case command == "!addskill":
            AddSkill(session, message)
        case command == "!ping":
            session.ChannelMessageSend(message.ChannelID, "Pong!")
        case command == "!help":
            session.ChannelMessageSend(message.ChannelID, "List of commands:\n`!addskill` - Start adding a new skill.\n`!ping` - Responds with 'Pong!'.\n`!help` - Shows this message.")
        default:
            session.ChannelMessageSend(message.ChannelID, "Unknown command. Type !help for a list of commands.")
        }
    }
}
