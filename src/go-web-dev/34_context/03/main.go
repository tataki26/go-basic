package main

import (
	"fmt"
	"time"
)

func main() {
	// n --> channel int
	for n := range gen() {
		fmt.Println(n)

		if n == 5 {
			break
		}
	}
	// after break, it sleeps
	time.Sleep(1 * time.Minute)
}

// gen is a broken generator that will leak a goroutine
// because go routine runs even if there is no input
// it needs signal ---> "we don't need what you sends to us"
// gen returns int channel
func gen() <-chan int {
	// create channel
	ch := make(chan int)

	// create go routine
	go func() {
		var n int

		for {
			ch <- n
			n++
		}
	}()

	return ch
}
