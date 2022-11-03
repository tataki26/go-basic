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

func handle(conn net.Conn) {
	// get pointer of Scanner
	// to use Scan method and Text method
	scanner := bufio.NewScanner(conn)
	// return boolean
	// end of input -> false
	for scanner.Scan() {
		// get the text
		ln := scanner.Text()
		fmt.Println(ln)
	}
	// we never get here (only when browser closes first, the codes below is called)
	// we have an open stream connection
	// how does the above reader know when it's done?
	defer conn.Close()

	fmt.Println("Code got here")
}
