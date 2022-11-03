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
	// .Listen(network, address)
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

		// .WriteString(Writer, string)
		// write string to Writer
		io.WriteString(conn, "\nHello from TCP server\n")
		// 지정한 출력 stream(conn)에 출력
		// it internally calls io.WriteString
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()

	}
}
