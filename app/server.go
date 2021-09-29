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
	_, err = l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()
}
