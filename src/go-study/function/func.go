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

	f1(3)

	s := sum(10, 0)
	fmt.Println(s)

	rtn := fib(10)
	fmt.Println(rtn)
}

// recursive function
func f1(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("f1(%d) before call f1(%d)\n", x, x-1)
	f1(x - 1)
	fmt.Printf("f1(%d) after call f1(%d)\n", x, x-1)
}

func sum(x, y int) int {
	if x == 0 {
		return y
	}
	y += x
	return sum(x-1, y)
}

func fib(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}
