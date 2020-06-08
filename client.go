package main

import (
	"fmt"
	"net"
)

func main() {
	mySocket, err := net.Dial("tcp", "10.1.1.27:50030")
	if err != nil {
		fmt.Printf("Dial error: %s\n", err)
		return
	}
	defer mySocket.Close()

	sendMsg := "Hello! This is test message from my sample client!"
	mySocket.Write([]byte(sendMsg + "\n"))
}
