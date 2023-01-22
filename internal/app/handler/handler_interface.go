package handler

import (
	"github.com/gorilla/mux"
)

type service interface {
	authenticationHandler()
	userMGMTHandler()
	homePageHandler()
}

type handler struct {
	router *mux.Router
}
