package main

import "fmt"

func main() {
	var i int
	var j int
	for i = 1; i < 10; i++ {
		fmt.Printf("%d단\n", i)
		for j = 1; j < 10; j++ {
			if i == 3 && j == 2 {
				continue
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}
