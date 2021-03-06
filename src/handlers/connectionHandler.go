package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Connect() {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	// if unauthorized, go get new token from website
	BotID = u.ID

	// important: add handler functions before .Open()
	// dg.AddHandler(messageHandler)
	dg.AddHandler(Ping)
	dg.AddHandler(CreationCommand)

	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
