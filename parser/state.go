package parser

import (
	"github.com/MattheDev53/GoMuteUs/discord"
)

func aliveParser(cmd string) {
	if cmd == "" {
		discord.ListNotInState(discord.Offline)
		return
	}
	switch cmd[0] {
	case '0':
		discord.MuteState(-1)
	case '1':
		discord.SetUserState(-1, discord.Alive)
	}
}

func deadParser(cmd string) {
	if cmd == "" {
		discord.ListInState(discord.Dead)
		return
	}
	switch cmd[0] {
	case '0':
		discord.ClearDead()
	case '1':
		discord.SetUserState(-1, discord.Dead)
	}
}

func offlineParser(cmd string) {
	if cmd == "" {
		discord.ListInState(discord.Offline)
		return
	}
	switch cmd[0] {
	case '0':
		discord.OptimizeUsers()
	case '1':
		discord.SetUserState(-1, discord.Offline)
	}
}
