package parser

import (
	"fmt"
	"github.com/MattheDev53/GoMuteUs/discord"
)

func muteParser(cmd string) {
	switch cmd[0] {
	case '0':
		fmt.Println("Mute Alive")
	case '1':
		discord.MuteUser(-1)
	}
}
