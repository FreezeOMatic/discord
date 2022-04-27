package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const token string = "OTY4ODMxMzAzNjg2NDIyNTg4.Ymkkdg.DkHP-KX7XEF0mIsFNIYzx1uiByI"

var BotID string

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running now ...")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session m *discordgo.MessageCreate) {
	fmt.Println(m.Content)
}
