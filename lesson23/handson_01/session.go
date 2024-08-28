package main

import "net/http"

func getUser(r *http.Request) user {
	var u user

	c, err := r.Cookie("session")
	if err != nil {
		retrun u
	}

	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return
}