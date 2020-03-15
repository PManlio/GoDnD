package main

import (
	"fmt"

	"./handlers"
)

// const token string = "Njg4NDMyMDUyMjY5MjE5OTMw.Xm0PaQ.fm_Is5nrNRCIYl3PBl5BZG6fMVo"

// var botID string

func main() {
	fmt.Println("main is actually working.")

	handlers.Connect()

	fmt.Println("bot is running.")

	<-make(chan struct{})
	return
}
