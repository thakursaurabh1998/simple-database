package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const (
	HOST = "localhost"
	PORT = "6379"
	TYPE = "tcp4"
)

func StartTCPServer() error {
	server, err := net.Listen(TYPE, fmt.Sprintf("%s:%s", HOST, PORT))
	defer server.Close()

	if err != nil {
		return fmt.Errorf("Error in starting the TCP server: %w", err)
	}

	for {
		conn, err := server.Accept()

		if err != nil {
			return fmt.Errorf("Error in creating connection: %w", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(client net.Conn) {
	defer client.Close()

	log.Printf("Serving %s\n", client.RemoteAddr().String())

	for {
		netData, err := bufio.NewReader(client).ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("Connection closed for client %s\n", client.RemoteAddr().String())
			} else {
				log.Println("Error in reading data from connection", err)
			}
			break
		}

		temp := strings.TrimSpace(netData)

		log.Println("Received: ", temp)

		time := time.Now().Format("2006-01-02 15:04:05")
		client.Write([]byte(time + " Received : " + temp + "\n"))
	}
}
