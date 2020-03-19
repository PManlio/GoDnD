package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../character"
)

func SaveToJsonFile(chara *character.Character) {
	// when we try to convert a file in Json, it is actually
	// converted into a byte sequence

	byteArray, err := json.Marshal(chara)
	// if we want a "more legible" version of the json object
	// we can use json.MarshalIndent
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

	_ = ioutil.WriteFile("characters.json", byteArray, 0644)
	/*
		ioutil doc:
		If the file does not exist, WriteFile creates it with
		permissions perm (before umask); otherwise WriteFile
		truncates it before writing
	*/
}
