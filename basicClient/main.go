package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	port := "3030"
	// Connect to server using tcp protocol
	connection, err := net.Dial("tcp", "localhost:"+port)

	if err != nil {
		log.Fatalln("Error connecting to remote server on port: "+port, err)
	}

	data := "test message"
	sendData(connection, data)
}

func sendData(connection net.Conn, data string) {

	//Add bytes to end with \n
	message := append([]byte(data), 13, 10)
	//write data to server
	sb, err := connection.Write(message)

	if err != nil {
		fmt.Println("Error writing to connection: " + err.Error())
	}

	fmt.Printf("%v bytes written to server\n", sb)

	defer exit(connection)
}

func exit(connection net.Conn) {
	//Before closing send exit  command
	//TODO check why this is not beem received correctly on the server side
	sb, err := connection.Write([]byte{13, 10})
	fmt.Printf("%v bytes written to server\n", sb)
	if err != nil {
		fmt.Println("Error writing to connection: " + err.Error())
	}
	defer connection.Close()
}
