package main

import (
	"bufio"
	"fmt"
	"github.com/escalopa/go-client-server-socket/constants"
	"net"
	"os"
)

func main() {
	// Read server ipaddress
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter server ipaddress:")
	ip, _ := reader.ReadString('\n')

	// Establish `connection` wit the server
	connection, err := net.Dial(constants.ServerType, ip[:len(ip)-1]+":"+constants.ServerPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	// Close `connection` at the end of the function
	defer connection.Close()

	// Read `buffer` sent from the server
	buffer := make([]byte, 1024)
	ml, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving response:", err.Error())
		os.Exit(1)
	}

	// Print message sent from the server, `ml` is the size of message in byes
	fmt.Println("Server response:", string(buffer[:ml]))
}
