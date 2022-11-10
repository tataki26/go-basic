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
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	i := 0
	var m, u string

	// read
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request method
			m = strings.Fields(ln)[0]
			// request URL
			u = strings.Fields(ln)[1]
			fmt.Println("Method:", m)
			fmt.Println("URL:", u)
		}
		if ln == "" {
			fmt.Println("This is the end of the HTTP Request headers")
			break
		}
		i++
	}

	// write
	body := "Method: "
	body += m
	body += "\n"
	body += "URL: "
	body += u

	// respond header in network tab
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	// body is displayed in the brower
	io.WriteString(conn, body)
}
