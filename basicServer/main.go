package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	//create listener using tcp
	port := "3030"
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalln("Error creating tcp listener on port: "+port, err)
	}

	fmt.Printf("Server Started. Waiting for new connection on port: %v \n", port)

	for {
		//Keeps waiting for new connections
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection on port: "+port, err)
		}

		//If no errors handle connection inside a new goroutine ( concurrently!)

		go handleConn(connection)
	}
}

func handleConn(connection net.Conn) {
	fmt.Println("New Connection accepted from: ", connection.RemoteAddr())

	//Copy connection buffer to console (NOT BEST PRACTICE, ONLY FOR LEARNING PURPOSE)
	//io.Copy(os.Stdout, connection)

	//Read buffer until newline firs occurs
	//TODO: review how to declare delimiter \n as byte outside of Readbytes func
	readFromConn(connection)
	//Close connection on func exit
	defer connection.Close()
	defer fmt.Printf("Connection: %v closed \n", connection.RemoteAddr())

}

func readFromConn(connection net.Conn) error {
	for {
		line, err := bufio.NewReader(connection).ReadBytes('\n')
		//Check specific error when end of file occurs ( exit)
		if err == io.EOF {
			//Log EOF
			fmt.Printf("EOF err reading from: %v\n", connection.RemoteAddr())
			return nil
		}
		//if any other error --> exit
		if err != nil {
			log.Fatalf("Error readin buffer from connection: %v \nwith error: %v", connection.RemoteAddr(), err)
		}
		//if Exit character found exit for loop
		exitChar := []byte{13, 10} //[13 10] =="\n"
		// fmt.Printf("exitChar: %v", exitChar)
		if bytes.Equal(exitChar, line) {
			fmt.Printf("Exit command %v received \n", exitChar)
			return nil
		}
		//Print line received from connection (Delete last 2 bytes (delimiter bytes))
		fmt.Printf("Conn: %v bytes: %v  string: %v\n", connection.RemoteAddr(), line, string(line[:len(line)-2]))

	}
}
