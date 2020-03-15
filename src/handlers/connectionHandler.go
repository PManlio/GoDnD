package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// get this from environment.yaml
//const token string = "Njg4NDMyMDUyMjY5MjE5OTMw.Xm0PaQ.fm_Is5nrNRCIYl3PBl5BZG6fMVo"
const token string = "Njg4NDMyMDUyMjY5MjE5OTMw.Xm1_rQ.Dhi2CMjah1x0gGLf33EI7q7ISpc"

var BotID string

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

	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
