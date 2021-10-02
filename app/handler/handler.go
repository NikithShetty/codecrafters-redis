package handler

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"

	"codecrafters-redis/app/commands"
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

		// in := []byte(cmd)
		inCmd := &commands.Command{CmdType: commands.PING, Data: []string{}}
		outChan := make(chan *commands.CmdResult)
		// Add cmd to be processed onto cmdQ
		cmdQ <- eventloop.CmdStruct{In: inCmd, Out: outChan}

		// Wait for response from event loop
		cmdRes := <-outChan
		close(outChan)

		res, err := encodeCmdResult(cmdRes)
		utils.CheckError(err)

		conn.Write(res)
	}
}

func encodeCmdResult(cmdRes *commands.CmdResult) ([]byte, error) {
	switch cmdRes.DataType {
	case commands.Error:
		res, err := cmdRes.GetErr()
		if err != nil {
			utils.CheckError(err)
		}

		return utils.ErrRespByte(res), nil
	case commands.SimpleString:
		res, err := cmdRes.GetSimpleStr()
		if err != nil {
			utils.CheckError(err)
		}

		return utils.OkRespByte(res), nil
	}

	return []byte{}, errors.New("Failed to encode datatype" + string(cmdRes.DataType))
}
