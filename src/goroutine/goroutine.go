package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 함수 시작", time.Now())

	// go routine 생성
	// main 함수가 종료되면 go routine 함께 종료
	go long()
	go short()

	// 5초 대기
	time.Sleep(5 * time.Second)
	fmt.Println("main 함수 종료", time.Now())
}

func long() {
	fmt.Println("long 함수 시작", time.Now())
	// 3초 대기
	time.Sleep(3 * time.Second)
	fmt.Println("long 함수 종료", time.Now())
}

func short() {
	fmt.Println("short 함수 시작", time.Now())
	// 1초 대기
	time.Sleep(1 * time.Second)
	fmt.Println("short 함수 종료", time.Now())
}
