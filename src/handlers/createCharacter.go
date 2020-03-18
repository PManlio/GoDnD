package handlers

import (
	"../character"
	"../utils"

	"github.com/bwmarrin/discordgo"

	"fmt"
)

var player = new(character.Character)

var warningString = ` wait, it's not your turn!
Please wait until your ` + player.PlayerName + ` ends its character creation`

func CreationCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

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
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandler(setName)
	}

	player.CharName = m.Content
	_, _ = s.ChannelMessageSend(ChannelID, "Good, now write your class")
	s.AddHandler(setClass)
	fmt.Println(player)
}

func setClass(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandler(setClass)
	}

	// here we take in input whatever the duck player insert the duo {class level}
	var newClass character.Class
	newClass.ClassName, newClass.Level = utils.ChatClassParser(m.Content)

	player.Classes = append(player.Classes, newClass)

	_, _ = s.ChannelMessageSend(ChannelID, `Great. Now move on and write your abilities
	(use this format, writing only your ability numbers: Stre, Dext, Cons, Inte, Wisd, Char)
	Example: 14, 18, 17, 10, 10, 12`)

	s.AddHandler(setAbilities)
	fmt.Println(player)
}

func setAbilities(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandler(setAbilities)
	}

	var ability character.Abilities
	var errmsg string
	ability.Stre, ability.Dext, ability.Cons, ability.Inte, ability.Wisd, ability.Char, errmsg = utils.ChatAbilitiesParser(m.Content)

	if errmsg != "" {
		_, _ = s.ChannelMessageSend(ChannelID, errmsg+"\nPlease try again.")
		s.AddHandler(setAbilities)
	}
	player.Abilities = ability
	fmt.Println(player) // SONO ARRIVATO QUI, PROVAMI
}

func setCompetence(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandler(setCompetence)
	}
}
func setHitpoints(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandler(setHitpoints)
	}
}
func setArmorClass(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandler(setArmorClass)
	}
}
func setInventory(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandler(setInventory)
	}
}
func saveChar(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandler(saveChar)
	}
}
