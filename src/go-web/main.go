package main

import (
	"go-web/myapp"
	"net/http"
)

func main() {
	// 웹 서버를 구동하고 웹 요청을 (대기하고) 받아 (미리 지정한 핸들러에 따라) 처리
	// 웹 요청을 처리할 핸들러를 전달하지 않으면 http.DefaultServeMux가 동작
	// 3000번 포트로 웹 서버 구동
	// myapp.NewHttpHandler의 return: mux
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
