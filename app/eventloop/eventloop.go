package eventloop

import (
	"codecrafters-redis/app/utils"
	"fmt"
)

func StartEventLoop(cmdQ ReadCmdQ) {
	fmt.Println("Started event loop")
	for cmdChan := range cmdQ {
		fmt.Println("Event loop: Got event")
		cmd := cmdChan.In

		fmt.Println("Event loop: ", string(cmd))
		// if strings.Contains(cmd, "PING") || strings.Contains(cmd, "ping") {
		// msg := cmd[5:]

		// fmt.Println("msg", msg, len(msg))
		// if len(msg) > 0 {
		// 	conn.Write(utils.okRespByte(msg))
		// } else {
		cmdChan.Out <- utils.OkRespByte("PONG")
		// }
		// } else {
		// 	conn.Write(errRespByte("-"))
		// }
	}
}
