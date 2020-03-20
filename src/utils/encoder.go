package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../character"
)

const charDir string = "/home/manlio/Scrivania"

func SaveToJsonFile(chara *character.Character, playerId string) {
	// when we try to convert a file in Json, it is actually
	// converted into a byte sequence

	byteArray, err := json.MarshalIndent(chara, " ", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

	ioutil.WriteFile(charDir+"/"+playerId+".json", byteArray, 0644)
}
