package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	greeting := "Hello, World!"
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

	// 새로운 파일 생성
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error: creating file", err)
	}
	// 지연: 페이지가 모두 로드된 후에 외부 스크립트 실행됨을 명시
	// close 지연 처리
	defer nf.Close()

	// strtpl 가져오기
	// 입출력을 위한 인터페이스 Reader/Writer
	// reader 인스턴스 생성
	// 문자열을 입력하여 새 파일에 작성
	io.Copy(nf, strings.NewReader(strtpl))
}
