package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	cfg "github.com/MattheDev53/GoMuteUs/config"
	"github.com/MattheDev53/GoMuteUs/parser"
	"github.com/MattheDev53/GoMuteUs/discord"

	"fmt"
)


func main() {
	discord.Conf = cfg.ReadConfig()
	var err error

	discord.Session, err = discordgo.New(discord.Conf.Token)
	if err != nil {
		log.Fatalf("Error Initializing: %v\n", err)
	}

	err = discord.Session.Open()
	if err != nil {
		log.Fatalf("Error Opening: %v\n", err)
	}
	defer discord.Session.Close()


	cmd := ""
	for cmd != "exit" {
		fmt.Printf("GoMuteUs |> ")
		fmt.Scanln(&cmd)
		parser.Parse(cmd)
	}
}

