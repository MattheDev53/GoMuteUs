package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	cfg "github.com/MattheDev53/GoMuteUs/config"

	"fmt"
)

type State int

const (
	Alive State = iota
	Dead
	Offline
)

var Members     []*discordgo.Member
var MemberState [99]State

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

	Members, err = discord.GuildMembers(conf.GuildID, "0", 99)
	if err != nil {
		log.Fatalf("Error Getting Guild Members: %v\n", err)
	}

	InitializeUserState()
	ListMembers(-1)

	cmd := ""
	for cmd != "exit" {
		fmt.Printf("GoMuteUs |> ")
		fmt.Scanln(&cmd)
	}
}

func ListMembers(state State) {
	for i := 0; i < len(Members); i++ {
		if MemberState[i] == state || state == -1 {
			fmt.Printf("%02d: %32s [%s]\n", i, Members[i].User.Username, StateName(MemberState[i]))
		}
	}
}

func InitializeUserState() {
	for i := 0; i < len(Members); i++ {
		MemberState[i] = Alive
	}
}

func StateName(s State) string {
	switch s {
	case Alive:
		return "LIVE"
	case Dead:
		return "DEAD"
	case Offline:
		return "AWAY"
	}
	return "UNKN"
}
