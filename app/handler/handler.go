package handler

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"codecrafters-redis/app/eventloop"
	"codecrafters-redis/app/utils"
)

func HandleConn(conn net.Conn, cmdQ eventloop.WriteCmdQ) {
	defer conn.Close()

	in := bufio.NewReader(conn)

	for {
		cmd, err := in.ReadString('\n')
		if err == io.EOF {
			break
		}
		utils.CheckError(err)

		fmt.Println("Read", len(cmd), "bytes")
		fmt.Println(cmd)

		in := []byte(cmd)
		outChan := make(chan []byte)
		// Add cmd to be processed onto cmdQ
		cmdQ <- eventloop.CmdStruct{In: in, Out: outChan}

		cmdRes := <-outChan
		close(outChan)

		conn.Write(cmdRes)
	}
}
