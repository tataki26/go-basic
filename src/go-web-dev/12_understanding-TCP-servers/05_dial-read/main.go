package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// client side
	// connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// read data from server
	bs, err := io.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	// convert bytes to string
	fmt.Println(string(bs))

}
