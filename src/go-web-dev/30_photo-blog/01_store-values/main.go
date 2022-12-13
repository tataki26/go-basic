package main

import (
	"net/http"
	"strings"
	"text/template"

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
	c = appendValue(w, c)

	// split to get each name of photos
	// it returns string slice
	xs := strings.Split(c.Value, "|")

	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

// get uuid at least
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	// no cookie --> create cookie
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

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// values
	photos := []string{"gangneung.jpg", "yeosu.jpg", "geoje.jpg"}

	// append
	// put the value of cookie(uuid) into s
	s := c.Value

	for _, p := range photos {
		// Contains: to report whether substring(p) is within string(s)
		// no p in s --> add p to s with "|"
		if !strings.Contains(s, p) {
			s += "|" + p
		}
	}

	// update the value of cookie(uuid+photos)
	c.Value = s

	http.SetCookie(w, c)

	return c
}
