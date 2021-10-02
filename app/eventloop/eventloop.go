package eventloop

import (
	"codecrafters-redis/app/commands"
	"fmt"
	"strings"
)

func StartEventLoop(cmdQ ReadCmdQ) {
	fmt.Println("Started event loop")
	for cmdChan := range cmdQ {
		cmd := cmdChan.In

		var res *commands.CmdResult

		fmt.Println("Event loop: Got command ", string(cmd.CmdType))

		switch cmd.CmdType {
		case commands.ECHO:
			res = echo(cmd.Data)

		case commands.PING:
			res = ping(cmd.Data)

		default:
			res = commands.NewErrResult("unknown command " + string(cmd.CmdType))
		}

		cmdChan.Out <- res
	}
}

func echo(str []string) *commands.CmdResult {
	return commands.NewBulkStringResult(strings.Join(str, " "))
}

func ping(str []string) *commands.CmdResult {
	if len(str) > 0 {
		return commands.NewBulkStringResult(strings.Join(str, " "))
	} else {
		return commands.NewBulkStringResult("PONG")
	}
}
