package main

import (
	"flag"
	"fmt"
	"net"

	"codecrafters-redis/app/eventloop"
	"codecrafters-redis/app/handler"
	"codecrafters-redis/app/utils"
)

func main() {
	fmt.Println("Starting tcp server..")

	port := flag.String("Port number", "6379", "Port number to which redis server should bind to.")
	addr := "0.0.0.0:" + *port

	l, err := net.Listen("tcp", addr)
	utils.CheckError(err)
	defer l.Close()

	fmt.Println("Listening at", addr)

	cmdQ := make(eventloop.CmdQ, 1)

	go eventloop.StartEventLoop(cmdQ)

	for {
		conn, err := l.Accept()
		utils.CheckError(err)

		go handler.HandleConn(conn, cmdQ)
	}
}
