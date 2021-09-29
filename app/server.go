package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting tcp server..")

	port := flag.String("Port number", "6379", "Port number to which redis server should bind to.")
	addr := "0.0.0.0:" + *port

	l, err := net.Listen("tcp", addr)
	checkError(err)
	defer l.Close()

	fmt.Println("Listening at", addr)

	for {
		conn, err := l.Accept()
		checkError(err)

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	in := bufio.NewReader(conn)

	cmd, err := in.ReadString('\n')
	checkError(err)

	fmt.Println("Read", len(cmd), "bytes")
	fmt.Println(cmd)

	if strings.Contains(cmd, "PING") {
		msg := cmd[5:]

		fmt.Println("msg", msg, len(msg))
		if len(msg) > 0 {
			conn.Write(okRespByte(msg))
		} else {
			conn.Write(okRespByte("PONG"))
		}
	} else {
		conn.Write(errRespByte("Unknown cmd"))
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
