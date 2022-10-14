package main

import "fmt"

func main() {
	var A [10]int
	for i := 0; i < len(A); i++ {
		A[i] = i * i
	}
	fmt.Println(A)

	s := "Hello World"
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), ", ")
	}
	fmt.Println()

	kor := "월드" // 3 byte
	fmt.Println("len(kor) =", len(kor))
	for i := 0; i < len(kor); i++ {
		fmt.Print(kor[i], ", ")
	}
	fmt.Println()

	// utf-8을 표현하는 타입
	kor2 := []rune(kor)
	fmt.Println("len(kor2) =", len(kor2))
	for i := 0; i < len(kor2); i++ {
		fmt.Print(kor2[i], ", ")
	}
}
