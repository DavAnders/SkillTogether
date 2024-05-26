package commands

import (
	"github.com/bwmarrin/discordgo"
)

func FindFriends(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "What interest tags are relevant to you? Reply with tags.")
	// need to fill out
}

// Will need to think about managing state, if it there is a better way