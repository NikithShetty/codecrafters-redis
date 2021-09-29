package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting tcp server..")

	port := flag.String("Port number", "6379", "Port number to which redis server should bind to.")
	addr := "0.0.0.0:" + *port

	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	fmt.Println("Listening at", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	in := make([]byte, 10)
	readN, err := conn.Read(in)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Read", readN, "bytes")
	fmt.Println(string(in))
	conn.Write([]byte("+PONG\r\n"))

	conn.Close()
}
