package main

import (
	"net/http"
)

func getUser(req *http.Request) user {
	var u user

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exists already, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	// no cookie ---> not logged in
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	// get the user name
	un := dbSessions[c.Value]
	// ok means whether data exists in the map or not
	_, ok := dbUsers[un]

	return ok
}
