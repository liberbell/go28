package main

import (
	"go28/lesson27/handson_04/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
}
