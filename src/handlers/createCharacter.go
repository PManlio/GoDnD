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
		_, _ = s.ChannelMessageSend(ChannelID, `Ok, let's create a new character!
		First of all, insert your Character name:`)

		// save AuthorID as playerID
		player.PlayerID = m.Author.ID
		player.PlayerName = m.Author.Username

		s.AddHandlerOnce(setName)

		fmt.Println(player)
	}
}

func setName(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandlerOnce(setName)
		return
	}

	player.CharName = m.Content
	_, _ = s.ChannelMessageSend(ChannelID, `Good, now write your class and level
	Like: Warrior, 5`)
	s.AddHandlerOnce(setClass)

	fmt.Println(player)
}

func setClass(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandlerOnce(setClass)
		return
	}

	// here we take in input whatever the duck player insert the duo {class level}
	var newClass character.Class
	newClass.ClassName, newClass.Level = utils.ChatClassParser(m.Content)

	player.Classes = append(player.Classes, newClass)

	_, _ = s.ChannelMessageSend(ChannelID, `Great. Now move on and write your abilities
	(use this format, writing only your ability numbers: Stre, Dext, Cons, Inte, Wisd, Char)
	Example: 14, 18, 17, 10, 10, 12`)

	s.AddHandlerOnce(setAbilities)
	fmt.Println(player)
}

func setAbilities(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandlerOnce(setAbilities)
		return
	}

	var ability character.Abilities
	var errmsg string
	ability.Stre, ability.Dext, ability.Cons, ability.Inte, ability.Wisd, ability.Char, errmsg = utils.ChatAbilitiesParser(m.Content)

	if errmsg != "" {
		_, _ = s.ChannelMessageSend(ChannelID, errmsg+"\nPlease try again.")
		s.AddHandlerOnce(setAbilities)
		return
	}
	player.Ability = ability
	fmt.Println(player)
	_, _ = s.ChannelMessageSend(ChannelID, `We're almost done. Please set your Bonus Competence:
	Example: 1 (or 2 or 3... it has to be only a number)`)

	s.AddHandlerOnce(setCompetence)
}

func setCompetence(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandlerOnce(setCompetence)
		return
	}

	player.Competence = utils.ChatGetNumber(m.Content)

	_, _ = s.ChannelMessageSend(ChannelID, `Please set your Hit Points:
	Example: 25`)

	s.AddHandlerOnce(setHitpoints)
	fmt.Println(player)
}

func setHitpoints(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandlerOnce(setHitpoints)
		return
	}

	player.Hitpoints = utils.ChatGetNumber(m.Content)

	_, _ = s.ChannelMessageSend(ChannelID, `...And your armor class:
	Example: 18`)

	s.AddHandlerOnce(setArmorClass)
	fmt.Println(player)
}

func setArmorClass(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandlerOnce(setArmorClass)
		return
	}

	player.ArmorClass = utils.ChatGetNumber(m.Content)

	_, _ = s.ChannelMessageSend(ChannelID, `Now list your items (those in your inventory).
	If you want, you can skip this part by writing:
	!empty`)

	s.AddHandlerOnce(setInventory)
	fmt.Println(player)
}

func setInventory(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Author.ID != player.PlayerID {
		_, _ = s.ChannelMessageSend(ChannelID, m.Author.Username+warningString)
		s.AddHandlerOnce(setInventory)
		return
	}

	if m.Content == (BotPrefix + "empty") {
		_, _ = s.ChannelMessageSend(ChannelID, `Character saved, have fun!`)
		saveChar()
		fmt.Println(player)
		return
	}

	player.Inventory = utils.ChatGetInventory(m.Content)
	_, _ = s.ChannelMessageSend(ChannelID, `Character saved, have fun!`)

	saveChar()
}

func saveChar() {
	utils.SaveToJsonFile(player, player.PlayerID)
	resetPlayer()
}

func resetPlayer() {
	player = new(character.Character)
	fmt.Println(player)
}
