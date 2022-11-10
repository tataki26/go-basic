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

		go serve(conn)
	}

}

func serve(conn net.Conn) {
	defer conn.Close()

	i := 0
	var m, u string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m = strings.Fields(ln)[0]
			u = strings.Fields(ln)[1]
		}
		if ln == "" {
			fmt.Println("This is the end of the HTTP Request header")
			break
		}
		i++
	}

	if m == "GET" {
		if u == "/" {
			index(conn)
		}

		if u == "/apply" {
			apply(conn)
		}
	}

	if m == "POST" {
		if u == "/apply" {
			applyProcess(conn)
		}
	}

}

func index(conn net.Conn) {
	body := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charet="UTF-8">
			<title></title>
			</head>
			<body>
				<h1>THIS IS INDEX PAGE</h1>
				<a href="/apply">apply</a>
			</body>
			</html>
			`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}

func apply(conn net.Conn) {
	body := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charet="UTF-8">
			<title></title>
			</head>
			<body>
				<h1>THIS IS APPLY PAGE</h1>
				<form method="POST" action="/apply">
					<input type="submit" value="apply">
				</form>
			</body>
			</html>
			`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}

func applyProcess(conn net.Conn) {
	body := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charet="UTF-8">
			<title></title>
			</head>
			<body>
				<h1>APPLY PROCESSING...</h1>
			</body>
			</html>
			`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}
