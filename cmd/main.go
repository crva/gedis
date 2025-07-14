package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/crva/gedis/internal/protocol"
	"github.com/crva/gedis/internal/store"
)

func handleConnection(conn net.Conn, store *store.GedisStore) {
	defer conn.Close()
	reader := bufio.NewReader(conn) // Create a buffered reader to read from the connection

	for {
		line, err := reader.ReadString('\n') // Read a line from the connection
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		response := protocol.HandleCommand(line, store) // Process the command using the protocol package
		conn.Write([]byte(response + "\n"))
	}
}

func startServer(address string, store *store.GedisStore) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err) // Unable to start the TCP server
	}
	defer listener.Close()

	fmt.Println("Server started on", address)

	for {
		conn, err := listener.Accept() // Program is blocking here, waiting for a new connection
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, store)
	}
}

func main() {
	store := store.NewStore()
	address := "localhost:8080"
	startServer(address, store)
}
