package main

import (
	"fmt"
	"github.com/escalopa/go-client-server-socket/constants"
	"net"
	"os"
	"time"
)

func main() {

	server, err := net.Listen(constants.ServerType, constants.ServerHost+":"+constants.ServerPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer server.Close()

	fmt.Println("Listening on " + constants.ServerHost + ":" + constants.ServerPort)
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
