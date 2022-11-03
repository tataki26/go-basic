package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// client side
	// connect server
	// send or listen to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// read conn (written to server)
	bs, err := io.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}
