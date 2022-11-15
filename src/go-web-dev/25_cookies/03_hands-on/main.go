package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// request cookie
	c, err := req.Cookie("my-cookie")
	// if there is no cookie, set cookie default
	if err == http.ErrNoCookie {
		// cookie struct
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	// convert ASCII to int
	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}

	// counter
	count++
	// convert int to ASCII
	c.Value = strconv.Itoa(count)

	// set cookie
	http.SetCookie(w, c)

	io.WriteString(w, c.Value)
}
