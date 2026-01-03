package parser

func Parse(cmd string) {
	if cmd == "" {
		return
	}
	switch cmd[0] {
	case '0':
		muteParser(cmd[1:])
	case '1':
		unmuteParser(cmd[1:])
	case '7':
		offlineParser(cmd[1:])
	case '8':
		aliveParser(cmd[1:])
	case '9':
		deadParser(cmd[1:])
	case '+':
		Parse(cmd[1:])
	}
}
