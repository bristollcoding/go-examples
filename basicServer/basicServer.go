package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func mainServer() {

	//create listener using tcp
	port := "3030"
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalln("Error creating tcp listener on port: ", port)
	}

	fmt.Printf("Server Started. Waiting for new connection on port: %v \n", port)

	for {
		//Keeps waiting for new connections
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection on port: ", port)
		}

		//If no errors handle connection inside a new goroutine ( concurrently!)

		go handleBasicConn(connection)
	}
}

func handleBasicConn(connection net.Conn) {
	fmt.Println("New Connection accepted from: ", connection.RemoteAddr())

	//Copy connection buffer to console (NOT BEST PRACTICE, ONLY FOR LEARNING PURPOSE)
	io.Copy(os.Stdout, connection)

	//Close connection on func exit
	defer connection.Close()
	defer fmt.Printf("Connection: %v closed \n", connection.RemoteAddr())

}
