package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)

	// same input, same output
	c = getCode("test@example.com")
	fmt.Println(c)

	// different input, different output
	c = getCode("test2@example.com")
	fmt.Println(c)
}

func getCode(s string) string {
	// hmac: hash-based message authentication code
	// we need hash to check whether users change the value stored in their machine or not
	// create new hash with hash type and key
	h := hmac.New(sha256.New, []byte("ourkey"))

	// write string to hash
	io.WriteString(h, s)

	// Sum: to append the current hash to byte slice and return it
	// assign hash as hexadecimal number
	return fmt.Sprintf("%x", h.Sum(nil))
}
