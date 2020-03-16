package handlers

import (
	"../character"

	"github.com/bwmarrin/discordgo"

	"fmt"
)

var player = new(character.Character)

func CreationCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	//player := new(character.Character)

	if m.Content == (BotPrefix + "create") {
		_, _ = s.ChannelMessageSend(ChannelID, "Ok, let's create a new character! \nFirst of all, insert your Character name:")

		// save AuthorID as playerID
		player.PlayerID = m.Author.ID
		player.PlayerName = m.Author.Username

		s.AddHandler(setName)

		fmt.Println(player)
	}
}

func setName(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+" wait, it's not your turn!\nPlease wait until "+player.PlayerName+" ends its character creation")
		s.AddHandler(setName)
	}
	player.CharName = m.Content
	_, _ = s.ChannelMessageSend(ChannelID, "Good, now write your class")
	s.AddHandler(setClass)
	fmt.Println(player)
}

func setClass(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+" wait, it's not your turn!\nPlease wait until "+player.PlayerName+" ends its character creation")
		s.AddHandler(setName)
	}
}
func setAbilities(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+" wait, it's not your turn!\nPlease wait until "+player.PlayerName+" ends its character creation")
		s.AddHandler(setAbilities)
	}
}
func setCompetence(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+" wait, it's not your turn!\nPlease wait until "+player.PlayerName+" ends its character creation")
		s.AddHandler(setCompetence)
	}
}
func setHitpoints(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+" wait, it's not your turn!\nPlease wait until "+player.PlayerName+" ends its character creation")
		s.AddHandler(setHitpoints)
	}
}
func setArmorClass(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+" wait, it's not your turn!\nPlease wait until "+player.PlayerName+" ends its character creation")
		s.AddHandler(setArmorClass)
	}
}
func setInventory(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+" wait, it's not your turn!\nPlease wait until "+player.PlayerName+" ends its character creation")
		s.AddHandler(setInventory)
	}
}
func saveChar(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+" wait, it's not your turn!\nPlease wait until "+player.PlayerName+" ends its character creation")
		s.AddHandler(saveChar)
	}
}
