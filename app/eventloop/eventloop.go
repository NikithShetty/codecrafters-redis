package eventloop

import (
	"codecrafters-redis/app/commands"
	"fmt"
)

func StartEventLoop(cmdQ ReadCmdQ) {
	fmt.Println("Started event loop")
	for cmdChan := range cmdQ {
		cmd := cmdChan.In

		fmt.Println("Event loop: Got command ", string(cmd.CmdType))
		// if strings.Contains(cmd, "PING") || strings.Contains(cmd, "ping") {
		// msg := cmd[5:]

		// fmt.Println("msg", msg, len(msg))
		// if len(msg) > 0 {
		// 	conn.Write(utils.okRespByte(msg))
		// } else {
		cmdChan.Out <- commands.NewSimpleStringResult("PONG")
		// }
		// } else {
		// 	conn.Write(errRespByte("-"))
		// }
	}
}
