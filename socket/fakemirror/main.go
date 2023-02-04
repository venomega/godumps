package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var (
	address string = "0.0.0.0:8080"
)

var logger *log.Logger = log.Default()

func handler(client net.Conn) {

}

func main() {

	socket, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer socket.Close()

	println("Listening on " + address)

	for {
		client, _ := socket.Accept()
		msg := fmt.Sprintf("connection from %v", client.LocalAddr())
		logger.Printf(msg)
		defer client.Close()

		buff := make([]byte, 3333)
		n, _ := client.Read(buff)

		fmt.Printf("%s", buff[:n])

		go handler(client)

	}

}
