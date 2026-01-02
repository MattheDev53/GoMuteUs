package parser

import (
	"fmt"
)

func unmuteParser(cmd string) {
	switch cmd[0] {
	case '0':
		fmt.Println("Mute Alive")
	case '1':
		fmt.Println("Mute User")
	}
}
