package user_menagment

import (
	"RegisterUser/repository/user"
	"net/http"
)

// Search User By ID
type searchUserByIDService interface {
	getUserById(*user.User) error
}

type searchUserByIDController struct {
	rw  http.ResponseWriter
	req *http.Request
}
