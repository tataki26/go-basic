package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

// packet handling function
// packet: formatted data block
func handle(conn net.Conn) {
	// get pointer of Scanner(bring data from os.Stind[console input])
	// to use Scan method and Text method
	scanner := bufio.NewScanner(conn)

	// return boolean
	// end of input -> false
	for scanner.Scan() {
		// get the text
		ln := scanner.Text()
		// print to server program
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never get here
	// not unitl the client ends first
	// we have an open stream connection [infinite for loop]
	// how does the above reader know when it's done?
	fmt.Println("Code got here")
}
