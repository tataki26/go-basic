package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// changed Id type to string
type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func StoreUsers(m map[string]User) {
	f, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// write to file with encoding map
	// convert into JSON
	json.NewEncoder(f).Encode(m)
}

func LoadUsers() map[string]User {
	m := make(map[string]User)

	f, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		// empty map
		return m
	}
	defer f.Close()

	// decode file data and put it into map
	// convert into map
	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}

	return m
}
