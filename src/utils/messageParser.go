package utils

import (
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile("[0-9]+")

func ChatParser(message string) (string, int) {
	var nospaces = strings.Replace(message, " ", "", -1)
	var number = regex.FindString(nospaces)

	var returnstr = strings.Replace(nospaces, number, "", -1)
	var returnnum, _ = strconv.Atoi(number)

	return returnstr, returnnum
}
