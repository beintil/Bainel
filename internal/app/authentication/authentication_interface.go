package authentication

import (
	"Bainel/repository/user"
	"net/http"
)

// Registration
type registerService interface {
	register() error
}

type registerController struct {
	rw           http.ResponseWriter
	req          *http.Request
	registerUser user.User
}

// Authorisation
type authorizationService interface {
	authorization() error
}

type authorizationController struct {
	rw    http.ResponseWriter
	req   *http.Request
	user  user.User
	login user.Login
}
