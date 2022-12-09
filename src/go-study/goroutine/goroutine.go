package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 함수 시작", time.Now())

	// channel 생성
	done := make(chan bool)

	// go routine 생성
	// main 함수가 종료되면 go routine 함께 종료
	go long(done)
	go short(done)

	// 5초 대기
	// time.Sleep(5 * time.Second)

	// channel로부터 값 수신
	// channel이 준비되지 않으면 준비될 때까지 대기
	// channel에 값이 들어오기 전에는 수신되지 않는다
	<-done
	<-done
	// done 채널로부터 메시지를 받아야 프로그램 종료
	fmt.Println("main 함수 종료", time.Now())
}

func long(done chan bool) {
	fmt.Println("long 함수 시작", time.Now())
	// 3초 대기
	time.Sleep(3 * time.Second)
	fmt.Println("long 함수 종료", time.Now())

	// channel에 값 전송
	// channel이 비어있지 않으면 전송되지 않는다
	done <- true
}

func short(done chan bool) {
	fmt.Println("short 함수 시작", time.Now())
	// 1초 대기
	time.Sleep(1 * time.Second)
	fmt.Println("short 함수 종료", time.Now())

	done <- true
}
