package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong"+m.ChannelID)
	}
}
