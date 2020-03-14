package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const token string = "Njg4NDMyMDUyMjY5MjE5OTMw.Xm0PaQ.fm_Is5nrNRCIYl3PBl5BZG6fMVo"

var botID string

func main() {
	fmt.Println("main is actually working.")

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	botID = u.ID

	// important: add handler functions before .Open()
	dg.AddHandler(messageHandler)

	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("bot is running.")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == botID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong"+m.ChannelID)
	}

	fmt.Println(m.Content)

}
