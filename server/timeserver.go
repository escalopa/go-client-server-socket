package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	server, err := net.Listen(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer server.Close()

	fmt.Println("Listening on " + ServerHost + ":" + ServerPort)
	fmt.Println("Waiting for client...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			os.Exit(1)
		}

		fmt.Println("client connected")
		go processClient(connection)
	}

}

func processClient(connection net.Conn) {
	_, err := connection.Write([]byte(time.Now().String()))
	if err != nil {
		fmt.Println("Error sending:", err)
	}

	defer connection.Close()
}
