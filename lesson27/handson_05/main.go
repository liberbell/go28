package main

import (
	"context"
	"go28/lesson27/handson_05/controllers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	access_url := "mongodb+srv://<go_user>:<go_userpassword>@cluster0.4zuzwrb.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0}"
	clientOptions := options.Client().ApplyURI(access_url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	uc := controllers.NewUserController(getSession())
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
	return s
}
