package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	// request cookie
	cookie, err := req.Cookie("session")
	// error means no cookie
	// if then, create new UUID
	if err != nil {
		id := uuid.New()
		cookie = &http.Cookie{
			Name: "session",
			// convert UUID into String
			Value: id.String(),

			// create only in HTTPS
			// send cookie only when browser and server communicate with HTTPS
			// Secure: true,

			// not able to access through JavaScript
			// send cookie only when browser sends HTTP request to server
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
