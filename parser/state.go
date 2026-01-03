package parser

import (
	"fmt"
	"github.com/MattheDev53/GoMuteUs/discord"
)

func aliveParser(cmd string) {
	if cmd == "" {
		fmt.Println("All Users:")
		discord.ListMembersNotInState(discord.Offline)
		return
	}
	switch cmd[0] {
	case '0':
		fmt.Println("Curiosity Killed The Cat")
	case '1':
		fmt.Println("Adding to Alive:")
		discord.SetUserState(-1, discord.Alive)
	case '+':
		fmt.Println("All Users:")
		discord.ListMembersNotInState(discord.Offline)
		Parse(cmd[1:])
	}
}

func deadParser(cmd string) {
	if cmd == "" {
		fmt.Println("Deadlist:")
		discord.ListMembersInState(discord.Dead)
		return
	}
	switch cmd[0] {
	case '0':
		fmt.Println("Clearing Deadlist")
		discord.ClearDead()
	case '1':
		fmt.Println("Adding to Deadlist")
		discord.SetUserDead(-1)
	case '+':
		fmt.Println("Deadlist:")
		discord.ListMembersInState(discord.Dead)
		Parse(cmd[1:])
	}
}

func offlineParser(cmd string) {
	if cmd == "" {
		fmt.Println("Offline Users:")
		discord.ListMembersInState(discord.Offline)
		return
	}
	switch cmd[0] {
	case '0':
		discord.OptimizeUsers()
	case '1':
		fmt.Println("Adding to Offline")
		discord.SetUserState(-1, discord.Offline)
	case '+':
		fmt.Println("Offline Users:")
		discord.ListMembersInState(discord.Offline)
		Parse(cmd[1:])
	}
}
