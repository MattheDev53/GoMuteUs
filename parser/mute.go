package parser

import (
	"fmt"
)

func muteParser(cmd string) {
	switch cmd[0] {
	case '0':
		fmt.Println("Mute Alive")
	case '1':
		fmt.Println("Mute User")
	}
}
