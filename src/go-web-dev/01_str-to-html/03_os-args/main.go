package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	greeting := os.Args[1]
	// 실행한 프로그램(경로)
	fmt.Println(os.Args[0])
	// 입력값 + HTML 파일
	fmt.Println(os.Args[1])
	strtpl := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>My First Go Web Page</title>
			</head>
			<body>
				<h1>` + greeting + `</h1>
			</body>
		</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(strtpl))
}
