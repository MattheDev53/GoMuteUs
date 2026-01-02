package parser

import (
	"github.com/MattheDev53/GoMuteUs/discord"
)

func unmuteParser(cmd string) {
	if cmd == "" {
		discord.UnmuteState(discord.Alive)
		return
	}
	switch cmd[0] {
	case '0':
		discord.UnmuteState(-1)
	case '1':
		discord.UnmuteUser(-1)
	}
}
