package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type CmdStruct struct {
	in  []byte
	out chan []byte
}

type CmdQ = chan CmdStruct

type ReadCmdQ = <-chan CmdStruct
type WriteCmdQ = chan<- CmdStruct

func main() {
	fmt.Println("Starting tcp server..")

	port := flag.String("Port number", "6379", "Port number to which redis server should bind to.")
	addr := "0.0.0.0:" + *port

	l, err := net.Listen("tcp", addr)
	checkError(err)
	defer l.Close()

	fmt.Println("Listening at", addr)

	cmdQ := make(CmdQ, 1)

	go startEventLoop(cmdQ)

	for {
		conn, err := l.Accept()
		checkError(err)

		go handleConn(conn, cmdQ)
	}
}

func startEventLoop(cmdQ ReadCmdQ) {
	fmt.Println("Started event loop")
	for cmdChan := range cmdQ {
		fmt.Println("Event loop: Got event")
		cmd := cmdChan.in

		fmt.Println("Event loop: ", string(cmd))
		// if strings.Contains(cmd, "PING") || strings.Contains(cmd, "ping") {
		// msg := cmd[5:]

		// fmt.Println("msg", msg, len(msg))
		// if len(msg) > 0 {
		// 	conn.Write(okRespByte(msg))
		// } else {
		cmdChan.out <- okRespByte("PONG")
		// }
		// } else {
		// 	conn.Write(errRespByte("-"))
		// }
	}
}

func handleConn(conn net.Conn, cmdQ WriteCmdQ) {
	defer conn.Close()

	in := bufio.NewReader(conn)

	for {
		cmd, err := in.ReadString('\n')
		if err == io.EOF {
			break
		}
		checkError(err)

		fmt.Println("Read", len(cmd), "bytes")
		fmt.Println(cmd)

		in := []byte(cmd)
		outChan := make(chan []byte)
		// Add cmd to be processed onto cmdQ
		cmdQ <- CmdStruct{in, outChan}

		cmdRes := <-outChan
		close(outChan)

		conn.Write(cmdRes)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func okRespByte(s string) []byte {
	return []byte("+" + s + "\r\n")
}

func errRespByte(err string) []byte {
	return []byte("-Error " + err + "\r\n")
}
