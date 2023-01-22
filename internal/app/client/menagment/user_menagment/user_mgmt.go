package user_menagment

import (
	"Bainel/internal/app/database"
	"Bainel/repository/user"
	"context"
	"log"
	"net/http"
)

var (
	collection = database.Collection()
	ctx        = context.TODO()
	err        error
)

func UserSearchByID(w http.ResponseWriter, r *http.Request) {
	err = searchUserByIDService.getUserById(searchUserByIDController{
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
