package controllers

import (
	"net/http"

	"go28/lesson27/handson_04/models"

	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James",
		Gender: "Male",
		Age:    45,
		Id:     p.ByName("id"),
	}
}
