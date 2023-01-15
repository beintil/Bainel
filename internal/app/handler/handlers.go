package handler

import (
	"RegisterUser/internal/app/controller/user_controllers"
	"github.com/gorilla/mux"
)

func Handler(router *mux.Router) {
	router.HandleFunc("/users/{id}", user_controllers.GetUsers).Methods("GET")
	router.HandleFunc("/register/", user_controllers.RegisterUser).Methods("POST")
}
