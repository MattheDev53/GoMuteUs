package parser

func Parse(cmd string) {
	switch cmd[0]{
	case '0':
		muteParser(cmd[1:])
	case '1':
		unmuteParser(cmd[1:])
	}
}
