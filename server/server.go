package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func createListener() net.Listener {
	ln, err := net.Listen("tcp", "10.1.1.27:50030")
	if err != nil {
		log.Fatal(err)
	}

	return ln
}

func listenWorker(ln net.Listener) {
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}

func main() {
	log.Print("Launching server..")
	ln := createListener()
	listenWorker(ln)
}
