package authentication

import (
	"Bainel/repository/user"
	"net/http"
)

// Registration
type registerService interface {
	register(user.User) error
}

type registerController struct {
	rw  http.ResponseWriter
	req *http.Request
}

// Authorisation
type authorizationService interface {
	authorization(user.User, user.Login) error
}

type authorizationController struct {
	rw  http.ResponseWriter
	req *http.Request
}
