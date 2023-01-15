package user_controllers

import (
	"RegisterUser/internal/app/database"
	"RegisterUser/repository/user"
	"context"
	"log"
	"net/http"
)

var (
	collection = database.Collection()
	ctx        = context.TODO()
	err        error
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	err = service.userService(registerUserController{
		user: user.User{},
		rw:   w,
		req:  r,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		write, err := w.Write([]byte(err.Error()))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Panic(write, err)
		}
	}

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	err = service.userService(getUserByIDController{
		users: &user.User{},
		rw:    w,
		req:   r,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		write, err := w.Write([]byte(err.Error()))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Panic(write, err)
		}
	}

	w.WriteHeader(http.StatusOK)
}
