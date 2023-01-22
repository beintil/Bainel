package handler

import "github.com/gorilla/mux"

func Handler(router *mux.Router) {
	service.authenticationHandler(handler{
		router: router,
	})

	service.userMGMTHandler(handler{
		router: router,
	})

	service.homePageHandler(handler{
		router: router,
	})
}
