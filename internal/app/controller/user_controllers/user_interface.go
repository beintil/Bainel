package user_controllers

import (
	"RegisterUser/repository/user"
	"net/http"
)

// Registration
type registerUserService interface {
	register(user.User) error
}

type registerUserController struct {
	rw  http.ResponseWriter
	req *http.Request
}

// Getting By ID
type getUserByIDService interface {
	getUserById(*user.User) error
}

type getUserByIDController struct {
	rw  http.ResponseWriter
	req *http.Request
}
