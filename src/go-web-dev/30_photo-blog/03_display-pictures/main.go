package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	// add route to serve pictures
	// StripPrefix: to return a handler that serves HTTP requests by removing the given prefix from request URL's path
	// and invoke handler(second parameter)
	// FileServer: to return a handler that serves HTTP requests with the contents of the file system rooted at root
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		// create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()

		io.Copy(h, mf)

		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		path := filepath.Join(wd, "public", "pics", fname)

		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		// copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)

		// add file name to this user's cookie
		c = appendValue(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")

	// sliced cookie values to only send over images
	tpl.ExecuteTemplate(w, "index.gohtml", xs[1:])
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()

		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)
	}

	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value

	if !strings.Contains(s, fname) {
		s += "|" + fname
	}

	c.Value = s

	http.SetCookie(w, c)

	return c
}
