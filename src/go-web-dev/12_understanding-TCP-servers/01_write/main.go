package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// li: Listener interface(Accept, Close, Addr)
	// Close: after program is done, close listener
	// Accept: accept request(connect, error)
	// .Listen(network[protocol], address): wait for network connection
	// accept connection from all network interface
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		// accept request
		// connection: read(Reader) / write(Writer)
		// able to read and write through conn
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// all methods below are the way to write data to file
		// they are also the way to send the string to client
		// .WriteString(Writer, string)
		// write string to Writer
		io.WriteString(conn, "\nHello from TCP server\n")
		// print to assigned stream(conn)
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()

	}
}
