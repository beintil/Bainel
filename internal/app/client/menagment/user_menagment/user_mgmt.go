package user_menagment

import (
	"Bainel/internal/app/database"
	"Bainel/pkg/error_handler/client_errors"
	"Bainel/repository/user"
	"context"
	"net/http"
)

var (
	collection = database.Collection
	ctx        = context.TODO()
	err        error
)

func UserSearchByID(w http.ResponseWriter, r *http.Request) {
	err = searchUserByIDService.getUserById(searchUserByIDController{
		rw:   w,
		req:  r,
		user: &user.User{},
	})

	if err != nil {
		client_errors.ErrorPanic(w, err)
	}

	w.WriteHeader(http.StatusOK)
}
