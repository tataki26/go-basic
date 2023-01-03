package main

import (
	"encoding/json"
	"fmt"
)

type model struct {
	State   bool
	Picture []string
}

func main() {
	// no value
	m := model{}

	// empty data - false, []
	fmt.Println(m)

	// go struct to JSON
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	// empty slice = null(JavaScript)
	fmt.Println(string(bs))
}
