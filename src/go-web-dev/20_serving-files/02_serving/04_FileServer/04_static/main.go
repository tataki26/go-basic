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

	//*** log.Fatal wrapping***
	// ListenAndServe returns error
	// wrapping for catching error
	// log.Fatal: catch any type of error, print it to log file and call os.Exit(1)
	// 0: success, other: error
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))

	//*** http.Error() ***
	// Error replies to the request with the specified error message and HTTP code
	// ex:) http.Error(w, "file not found", 404)
}
