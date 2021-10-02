package handler

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"codecrafters-redis/app/commands"
	"codecrafters-redis/app/eventloop"
	"codecrafters-redis/app/utils"
)

func HandleConn(conn net.Conn, cmdQ eventloop.WriteCmdQ) {
	defer func() {
		conn.Close()

		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "fatal error: %s\n", err)
		}
	}()

	for {
		inCmd, err := processRedisCmd(conn)
		if err == io.EOF {
			break
		}
		utils.CheckError(err)

		// utils.LogInfo("Decoded command ", inCmd.CmdType, inCmd.Data)

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
	const RESPDELIM = "\r\n"

	switch cmdRes.DataType {
	case commands.Error:
		res, err := cmdRes.GetErr()
		if err != nil {
			return []byte{}, err
		}

		return utils.ErrRespByte(res), nil

	case commands.SimpleString:
		res, err := cmdRes.GetSimpleStr()
		if err != nil {
			return []byte{}, err
		}

		return utils.OkRespByte(res), nil

	case commands.BulkString:
		res, err := cmdRes.GetBlkStr()
		if err != nil {
			return []byte{}, err
		}

		if res == nil {
			str := "$-1" + RESPDELIM
			return []byte(str), nil
		} else {
			strLen := fmt.Sprint(len(*res))
			str := "$" + strLen + RESPDELIM + *res + RESPDELIM
			return []byte(str), nil
		}
	}

	return []byte{}, errors.New("failed to encode datatype " + string(cmdRes.DataType))
}

func processRedisCmd(conn net.Conn) (*commands.Command, error) {
	// utils.LogInfo("processRedisCmd")
	inStream := bufio.NewReader(conn)

	firstByte, err := inStream.ReadByte()
	utils.CheckError(err)

	var cmd *commands.Command

	switch firstByte {
	// Array
	case '*':
		arrLen, err := readLen(inStream)
		if err != nil {
			return nil, err
		}

		// utils.LogInfo("processRedisCmd arrLen", arrLen)

		// err = clearCLRF(inStream)
		// if err != nil {
		// 	return nil, err
		// }

		arr, err := readArray(arrLen, inStream)
		if err != nil {
			return nil, err
		}

		cmd, err = readCmd(arr)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("illegal char " + string(firstByte))
	}

	return cmd, nil
}

func readNBytes(n int, inStream *bufio.Reader) ([]byte, error) {
	// utils.LogInfo("readNBytes", n)
	var outBytes []byte
	for i := 0; i < n; i++ {
		b, err := inStream.ReadByte()
		if err != nil {
			return []byte{}, err
		}
		outBytes = append(outBytes, b)
	}

	return outBytes, nil
}

func readArray(n int, inStream *bufio.Reader) ([]string, error) {
	// utils.LogInfo("readArray", n)
	var outStr []string

	for i := 0; i < n; i++ {
		b, err := inStream.ReadByte()
		if err != nil {
			return []string{}, err
		}

		switch b {
		// Bulk string
		case '$':
			strLen, err := readLen(inStream)
			if err != nil {
				return []string{}, err
			}

			str, err := readString(strLen, inStream)
			if err != nil {
				return []string{}, err
			}

			outStr = append(outStr, str)

			err = clearCLRF(inStream)
			if err != nil {
				return nil, err
			}

		default:
			return []string{}, errors.New("illegal char " + string(b))
		}
	}

	return outStr, nil
}

func readString(n int, inStream *bufio.Reader) (string, error) {
	// utils.LogInfo("readString", n)
	outBytes, err := readNBytes(n, inStream)
	if err != nil {
		return "", err
	}

	return string(outBytes), nil
}

func readLen(inStream *bufio.Reader) (int, error) {
	// utils.LogInfo("readLen")
	var readLen = 0

	s, err := inStream.ReadString('\n')
	if err != nil {
		return -1, err
	}
	s = s[:len(s)-2]
	readLen, err = strconv.Atoi(s)
	if err != nil {
		return -1, err
	}

	return readLen, nil
}

func readCmd(arr []string) (*commands.Command, error) {
	// utils.LogInfo("readCmd", arr)
	cmdStr := arr[0]

	var cmd *commands.Command

	switch cmdStr {
	case string(commands.ECHO):
		cmd = &commands.Command{CmdType: commands.ECHO, Data: arr[1:]}

	case string(commands.PING):
		cmd = &commands.Command{CmdType: commands.PING, Data: arr[1:]}

	case string(commands.GET):
		cmd = &commands.Command{CmdType: commands.GET, Data: arr[1:]}

	case string(commands.SET):
		cmd = parseSetOptions(arr[1:])
		// cmd = &commands.Command{CmdType: commands.SET, Data: arr[1:]}

	default:
		return nil, errors.New("unknown command " + cmdStr)
	}

	return cmd, nil
}

func parseSetOptions(arr []string) *commands.Command {
	// utils.LogInfo("parseSetOptions", arr)
	expiryIndex := -1
	for ix, elem := range arr {
		if elem == "px" {
			expiryIndex = ix
			break
		}
	}

	// utils.LogInfo("parseSetOptions", expiryIndex)

	if expiryIndex == -1 {
		return &commands.Command{CmdType: commands.SET, Data: arr}
	} else {
		return &commands.Command{CmdType: commands.SETPX, Data: arr}
	}
}

func clearCLRF(inStream *bufio.Reader) error {
	cl, err := inStream.ReadByte()
	if err != nil {
		return err
	} else if cl != '\r' {
		return errors.New("clearCLRF: not '\\r' found " + string(cl) + " instead")
	}

	rf, err := inStream.ReadByte()
	if err != nil {
		return err
	} else if rf != '\n' {
		return errors.New("clearCLRF: not '\\n'" + string(rf) + " instead")
	}

	return nil
}
