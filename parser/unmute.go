package parser

import (
	"fmt"
	"github.com/MattheDev53/GoMuteUs/discord"
)

func unmuteParser(cmd string) {
	if cmd == "" {
		fmt.Println("Unmuting Alive:")
		discord.UnmuteState(discord.Alive)
		return
	}
	switch cmd[0] {
	case '0':
		fmt.Println("Unmuting Everyone:")
		discord.UnmuteState(-1)
	case '1':
		fmt.Println("Unmuting User:")
		discord.UnmuteUser(-1)
	}
}
