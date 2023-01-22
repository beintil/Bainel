package handler

import (
	"Bainel/internal/app/authentication"
	"Bainel/internal/app/client/menagment/user_menagment"
	"Bainel/internal/app/home"
)

func (hr handler) authenticationHandler() {
	// registration - method = "POST"
	hr.router.HandleFunc("/register/", authentication.Registration).Methods("POST")

	// authentication - method = "GET"
	hr.router.HandleFunc("/login/", authentication.Authorization).Methods("GET")
}

func (hr handler) userMGMTHandler() {
	// search user from id - method = "GET"
	hr.router.HandleFunc("/users/{id}", user_menagment.UserSearchByID).Methods("GET")
}

func (hr handler) homePageHandler() {
	// home page - method = "GET"
	hr.router.HandleFunc("/home/", home.Home).Methods("GET")
}
