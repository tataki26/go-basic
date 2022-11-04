package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	defer conn.Close()

	// instructions
	// write(print) to server program
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	// read & write
	// map[key]value
	// create memeory database
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// read(get data)
		ln := scanner.Text()
		// split the string around each instance of one or more consecutive white space character
		// and store it into the slice
		// data from user input
		fs := strings.Fields(ln)

		// logic
		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			k := fs[1] // key
			// get data from database by using index 1 element of user input as key
			v := data[k] // data[key] => value
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET":
			// not enough data
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE\r")
				continue
			}
			// data from user input
			k := fs[1]
			v := fs[2]
			// set data to database
			data[k] = v
		case "DEL":
			// delete data from database
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+fs[0]+"\r\n")
			continue
		}

	}
}
