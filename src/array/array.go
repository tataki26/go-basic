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
	fmt.Println()

	// 배열 복사
	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < 5; i++ {
		clone[i] = arr[i]
	}

	fmt.Println(clone)

	// 배열 역순
	temp := [5]int{}
	idx := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		temp[i] = arr[idx-i]
	}
	fmt.Println(temp)

	for i := 0; i < len(arr); i++ {
		arr[i] = temp[i]
	}
	fmt.Println(arr)

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println(arr)

}
