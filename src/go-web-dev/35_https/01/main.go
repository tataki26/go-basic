package main

import "net/http"

func main() {
	http.HandleFunc("/", foo)

	// TLS with cert file and key file
	// the cert file should be the concatenation of the server's certificate, any intermediates, and CA's certificate
	// 10443 for HTTPS server(443 port)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}
