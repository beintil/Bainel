package authentication

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

func RegisterHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./internal/app/authentication/register.html")
}

func Registration(w http.ResponseWriter, r *http.Request) {
	err = registerService.register(registerController{
		rw:  w,
		req: r,
	}, user.User{})

	if err != nil {
		client_errors.ErrorPanic(w, err)
	}
}

func Authorization(w http.ResponseWriter, r *http.Request) {
	err = authorizationService.authorization(authorizationController{
		rw:  w,
		req: r,
	}, user.User{}, user.Login{})

	if err != nil {
		client_errors.ErrorPanic(w, err)
	}

	w.WriteHeader(http.StatusOK)
}
