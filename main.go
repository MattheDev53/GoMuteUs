package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	cfg "github.com/MattheDev53/GoMuteUs/config"
)

func main() {
	conf := cfg.ReadConfig()

	discord, err := discordgo.New(conf.Token)
	if err != nil {
		log.Fatalf("Error Initializing: %v\n", err)
	}

	err = discord.Open()
	if err != nil {
		log.Fatalf("Error Opening: %v\n", err)
	}
	defer discord.Close()

	discord.ChannelMessageSend("1403578645489913858", "Ohayougozaimasu")
	discord.GuildMembers("1403578644487471227", "837074864845750272", 0)
	discord.ChannelMessageSend("1403578645489913858", "Muted????")
}
