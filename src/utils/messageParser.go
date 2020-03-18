package utils

import (
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile("[0-9]+")

func ChatClassParser(message string) (string, int) {
	var nospaces = strings.Replace(message, " ", "", -1)
	var number = regex.FindString(nospaces)

	var returnstr = strings.Replace(nospaces, number, "", -1)
	var returnnum, _ = strconv.Atoi(number)

	return returnstr, returnnum
}

func ChatAbilitiesParser(message string) (int, int, int, int, int, int, string) {
	var msgNumbers = regex.FindAllString(message, -1)
	var values []int
	for _, element := range msgNumbers {
		var number, _ = strconv.Atoi(element)
		values = append(values, number)
	}

	if len(values) > 6 {
		var errmsg = "TOO MANY ARGUMENTS!"
		return 0, 0, 0, 0, 0, 0, errmsg
	} else if len(values) < 6 {
		var errmsg = "SOME ARGUMENT IS MISSING!"
		return 0, 0, 0, 0, 0, 0, errmsg
	}

	var errmsg = ""
	return values[0], values[1], values[2], values[3], values[4], values[5], errmsg

}
