package throwHandlers

import(
	"../rolls"
	"../utils"
	"../character"

	"github.com/bwmarrin/discordgo"

	"fmt"
	"strings"
)

var userID

func attackCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == (BotPrefix + "attack") {
		userID = m.Author.ID
		_, _ = s.ChannelMessageSend(ChannelID, `Ok, tell me which stats gives you bonus:
		(example: dex / str / wis / int / car. Pick one)`)
		s.AddHandlerOnce(attackThrow)
	}
}

func attackThrow(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	player := &character.Character()
	player = utils.FindPlayer(userID)

	roll := rolls.AttackRoll()

	switch strings.ToLower(m.Content) {
	case "str":
		dicetype = 4
	case "dex":
		dicetype = 6
	case "int":
		dicetype = 8
	case "wis":
		dicetype = 10
	case "car":
		dicetype = 12
	}
}