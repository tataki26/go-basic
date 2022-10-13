package main

import "fmt"

func main() {
	// 1
	var a int = 3
	// 2
	var b int
	b = 4
	// 3
	c := 5

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("result = %v\n", b|c)
	fmt.Println("result2 =", c^a)

	fmt.Println(b << 1) // 100 << 1 // 1000 [8]
	fmt.Println(b >> 1) // 100 >> 1 // 10 [2]

	var d bool = 3 < 4 && 2 < 5

	fmt.Println(d)
}
