package main

import (
	"log"
	"net/http"
)

func main() {
	// *** static file server ***
	// serve all contents under current direcotry
	// as a special case, the returned file server redirects any request ending in "/index.html" to the same path,
	// without the final "index.html"
	// if we have index.html, Dir(.) doesnt' show all contents list in directory --> it shows web page (serve index.html)
	// redirection: a technique to give more than one URL address to a page/form/web app
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
