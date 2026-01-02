package parser

import (
	"fmt"
	"github.com/MattheDev53/GoMuteUs/discord"
)

func muteParser(cmd string) {
	if cmd == "" {
		fmt.Println("Muting Alive:")
		discord.MuteState(discord.Alive)
		return
	}
	switch cmd[0] {
	case '0':
		fmt.Println("Muting Everyone:")
		discord.MuteState(-1)
	case '1':
		fmt.Println("Muting User:")
		discord.MuteUser(-1)
	}
}
