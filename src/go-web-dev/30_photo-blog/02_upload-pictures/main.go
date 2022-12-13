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
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	// process form submission
	if req.Method == http.MethodPost {
		// request nf(new file) to the form
		// multipart.File is not same with File(os)
		// read whole files
		// fh: file header
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		// sha1 algorithm - sha1 hash(unique output - always same)
		// when multiple users upload file to server, unique identifier(file name) is needed
		// create sha for file name
		// split the file header
		// extract filename extension
		ext := strings.Split(fh.Filename, ".")[1]
		// return hash
		h := sha1.New()

		// copy from mf to h
		io.Copy(h, mf)

		// store result of Sum in fname
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// create new file
		// return current working directory
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		// create file path
		path := filepath.Join(wd, "public", "pics", fname)
		// create file(os)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		// Seek: to set the offset for the next Read or Wirte to offset
		// it returns the new offset relative to the start of the file or an error
		// since the offset was changed, the read/write header should be reset
		// SeekStart: relative to the start of the file
		// SeekCurrent: relative to the current offset
		// SeekEnd: relative to the end
		mf.Seek(0, 0)

		// copy the content of mf to file from os
		io.Copy(nf, mf)

		// add filename to this user's cookie
		c = appendValue(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")

	tpl.ExecuteTemplate(w, "index.gohtml", xs)
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

// takes in a file name now also
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value

	if !strings.Contains(s, fname) {
		s += "|" + fname
	}

	c.Value = s

	http.SetCookie(w, c)

	return c
}
