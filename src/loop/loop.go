package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println("최종 i 값:", i)
	var j int // scope of variable
	for j = 0; j < 10; j += 2 {
		fmt.Printf("%d ", j)
	}
	fmt.Println("최종 j 값:", j)
	var k int
	for {
		if k == 5 {
			k++
			continue
		}
		if k == 7 {
			break
		}
		fmt.Printf("%d ", k)
		k++
	}
}
