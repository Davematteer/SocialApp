package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Fprintf(conn, "Error reading command: %v\n", err)
		return
	}

	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)

	if len(parts) != 2 {
		fmt.Fprintf(conn, "Invalid command format. Expected format: COMMAND: RESOURCE \n ")
	}

}
