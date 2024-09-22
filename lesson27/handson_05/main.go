package main

import (
	"go28/lesson27/handson_05/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	access_url := "mongodb+srv://<go_user>:<go_userpassword>@cluster0.4zuzwrb.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0}"
	s, err := mgo.Dial(access_url)
	if err != nil {
		panic(err)
	}

}
