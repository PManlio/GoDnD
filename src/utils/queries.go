package utils

import (
	"encoding/json"
	"io/ioutil"

	"../character"
)

func checkerr(e error) {
	if e != nil {
		panic(e)
	}
}

func FindPlayer(playerId string) *character.Character {
	playerFile, err := ioutil.ReadFile(CharDir + "/" + playerId + ".json")
	checkerr(err)
	player := &character.Character{}
	err = json.Unmarshal(playerFile, player)
	checkerr(err)

	return player
}
