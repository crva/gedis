package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"

	"github.com/crva/gedis/internal/protocol"
	"github.com/crva/gedis/internal/store"
)

func handleConnection(conn net.Conn, store *store.GedisStore, aof *protocol.AOF) {
	defer conn.Close()
	reader := bufio.NewReader(conn) // Create a buffered reader to read from the connection

	for {
		line, err := reader.ReadString('\n') // Read a line from the connection
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		response := protocol.HandleCommand(line, store, aof) // Process the command using the protocol package
		conn.Write([]byte(response + "\n"))
	}
}

func startServer(address string, store *store.GedisStore, aof *protocol.AOF) {
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

		go handleConnection(conn, store, aof)
	}
}

func main() {
	replayAOF := flag.Bool("replay", false, "Replay AOF file on startup")
	host := flag.String("host", "localhost", "Host to bind the server")
	port := flag.Int("port", 8080, "Port to bind the server")
	flag.Parse()

	aof, err := protocol.NewAOF("gedis.aof")
	if err != nil {
		fmt.Println("Error creating AOF file:", err)
		return
	}
	defer aof.Close()

	store := store.NewStore()

	if *replayAOF {
		err := protocol.ReplayAOF("gedis.aof", store)
		if err != nil {
			fmt.Println("Error replaying AOF file:", err)
			return
		}
	}

	address := fmt.Sprintf("%s:%d", *host, *port)
	startServer(address, store, aof)
}
