package main

import (
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID(key), user
var dbSessions = map[string]string{} // session ID(key), user ID
// same code
// var dbSessions = make(map[string]string)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	// if there is no cookie, create new cookie
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	var u user
	// request the data which does not exist
	// return zero(0) and ok is false
	// if ok is true, bring user data from dbUsers with user name as key
	if un, ok := dbSessions[c.Value]; ok { // c.Value --> sID
		u = dbUsers[un]
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		u := user{un, f, l}
		dbSessions[c.Value] = un // sID as key
		dbUsers[un] = u          // user name as key
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[c.Value]
	// there is no user id to match c.Value
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un]

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
