package main

import (
	"go28/lesson27/handson_03/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="utf-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/111222">Go to: http://localhost:8080/user/111222</a>
	</body>
	</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}
}
