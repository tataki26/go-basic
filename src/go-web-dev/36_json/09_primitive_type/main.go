package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var a string

	rcvd := `"Todd"`
	err := json.Unmarshal([]byte(rcvd), &a)
	if err != nil {
		log.Fatalln(err)
	}

	var b int

	rcvd = `42`
	err = json.Unmarshal([]byte(rcvd), &b)
	if err != nil {
		log.Fatalln(err)
	}

	var c bool

	rcvd = `true`
	err = json.Unmarshal([]byte(rcvd), &c)
	if err != nil {
		log.Fatalln(err)
	}

	// null --> string --> empty string(= space) ( )
	// null --> []string --> empty slice ([])
	var d []string

	rcvd = `null`
	err = json.Unmarshal([]byte(rcvd), &d)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
}
