package eventloop

import (
	"fmt"
	"strings"

	"codecrafters-redis/app/commands"
	"codecrafters-redis/app/database"
)

func StartEventLoop(db *database.RedisStore, cmdQ ReadCmdQ) {
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

		case commands.GET:
			res = get(db, strings.Join(cmd.Data, " "))

		case commands.SET:
			res = set(db, cmd.Data[0], strings.Join(cmd.Data[1:], " "))

		default:
			res = commands.NewErrResult("unknown command " + string(cmd.CmdType))
		}

		cmdChan.Out <- res
	}
}

func CreateCmdQueue() CmdQ {
	return make(CmdQ, 1)
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

func get(db *database.RedisStore, key string) *commands.CmdResult {
	v := db.Get(key)
	return commands.NewBulkStringResult(v)
}

func set(db *database.RedisStore, key string, value string) *commands.CmdResult {
	db.Set(key, value)
	return commands.NewSimpleStringResult("OK")
}
