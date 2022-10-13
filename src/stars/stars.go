package main

import "fmt"

func main() {
	var i int
	var j int
	for i = 0; i < 5; i++ {
		for j = 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	for i = 0; i < 5; i++ {
		for j = 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i = 0; i < 4; i++ {
		for j = 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j = 0; j < i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i = 0; i < 3; i++ {
		for j = 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j = 0; j < 5-i*2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
