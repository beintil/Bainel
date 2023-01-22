package authentication

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

func Registration(w http.ResponseWriter, r *http.Request) {
	err = registerService.register(registerController{
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

func Authorization(w http.ResponseWriter, r *http.Request) {
	err = authorizationService.authorization(authorizationController{
		rw:  w,
		req: r,
	}, user.User{}, user.Login{})

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
