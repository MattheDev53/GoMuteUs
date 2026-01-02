package parser

import (
	"github.com/MattheDev53/GoMuteUs/discord"
)

func muteParser(cmd string) {
	if cmd == "" {
		discord.MuteState(discord.Alive)
		return
	}
	switch cmd[0] {
	case '0':
		discord.MuteState(-1)
	case '1':
		discord.MuteUser(-1)
	}
}
