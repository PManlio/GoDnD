package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../character"
)

// TODO: get CharDir from environment.yaml
const CharDir string = "/home/manlio/Scrivania"

func SaveToJsonFile(chara *character.Character, playerId string) {
	// when we try to convert a file in Json, it is actually
	// converted into a byte sequence

	byteArray, err := json.MarshalIndent(chara, " ", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

	ioutil.WriteFile(CharDir+"/"+playerId+".json", byteArray, 0644)
}
