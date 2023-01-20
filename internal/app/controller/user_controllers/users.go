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
	err = registerUserService.register(registerUserController{
		rw:  w,
		req: r,
	}, user.User{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		write, err := w.Write([]byte(err.Error()))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Panic(write, err)
		}
		return
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	err = getUserByIDService.getUserById(getUserByIDController{
		rw:  w,
		req: r,
	}, &user.User{})

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
