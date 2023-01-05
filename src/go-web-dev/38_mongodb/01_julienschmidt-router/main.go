package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// get the router
	r := httprouter.New()
	// use http method
	// when some request enters GET route("/"), handle it with index hanlder
	r.GET("/", index)

	// use custome router(httprouter) - not nil
	http.ListenAndServe("localhost:8080", r)
}

// handler of httprouter needs three parameter
// if you don't want to use it, use _
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
