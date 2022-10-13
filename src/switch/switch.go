package main

import "fmt"

func main() {
	x := 30

	// 변수 생략 시, true 값 찾음
	switch {
	case x > 40:
		fmt.Println("x는 40보다 크다")
	case x < 20:
		fmt.Println("x는 20보다 작다")
	// 참인 case가 여러 개이면 첫번째 case만 실행
	case x == 30:
		fmt.Println("x는 30이다")
	case x > 20:
		fmt.Println("x는 20보다 크다")
	}
}
