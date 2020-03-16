package handlers

import (
	"../character"

	"github.com/bwmarrin/discordgo"

	"fmt"
)

var player = character.Character{}

func CreationCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	//player := new(character.Character)

	if m.Content == (BotPrefix + "create") {
		_, _ = s.ChannelMessageSend(ChannelID, "Ok, let's create a new character! \nFirst of all, insert your Character name:")

		// save AuthorID as playerID
		player.PlayerID = m.Author.ID

		s.AddHandler(setName)

		fmt.Println(player)
	}
}

func setName(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != player.PlayerID {
		return
	}
	player.Name = m.Content
	fmt.Println(player)
}

func setClass(s *discordgo.Session, m *discordgo.MessageCreate)      {}
func setAbilities(s *discordgo.Session, m *discordgo.MessageCreate)  {}
func setHitpoints(s *discordgo.Session, m *discordgo.MessageCreate)  {}
func setArmorClass(s *discordgo.Session, m *discordgo.MessageCreate) {}
func setInventory(s *discordgo.Session, m *discordgo.MessageCreate)  {}
func saveChar(s *discordgo.Session, m *discordgo.MessageCreate)      {}
