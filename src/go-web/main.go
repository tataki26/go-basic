package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

// fooHandler의 메소드
// Handler 인터페이스의 ServeHTTP 메소드 구현
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "First page")
}

func main() {
	// 1
	// url별로 요청을 처리할 핸들러 함수 등록
	// "/" 경로(index page)로 접속했을 때 처리할 핸들러 함수 지정
	// http.ResponseWriter 타입에 값을 쓰면 HTTP 응답으로 전송
	// http.Request: 사용자가 요청한 정보
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// "Hello, Wolrd!"를 화면에 출력
		// Fprint: 지정한 스트림에 출력
		fmt.Fprintln(w, "Hello, World!")
	})

	// 1-1
	// 함수를 따로 정의하고 인자로 넘기기
	// 경로("/1")에 따른 핸들러
	http.HandleFunc("/1", FirstHandler)

	// 2
	// function이 아닌 instance를 등록
	http.Handle("/foo", &fooHandler{})

	// 웹 서버를 구동하고 웹 요청을 (대기하고) 받아 (미리 지정한 핸들러에 따라) 처리
	// 웹 요청을 처리할 핸들러를 전달하지 않으면 http.DefaultServeMux가 동작
	// 3000번 포트로 웹 서버 구동
	http.ListenAndServe(":3000", nil)
}
