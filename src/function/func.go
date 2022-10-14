package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func func1(x, y int) (int, int) {
	func2(x, y)
	return y, x
}

func func2(x, y int) {
	fmt.Println("func2", x, y)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

	a, b := func1(2, 3)
	fmt.Println(a, b)
}
