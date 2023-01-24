package myapp

import (
	"Bainel/internal/app/handler"
	"Bainel/pkg/error_handler/server_errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

var (
	router    = mux.NewRouter()
	lineError string
)

func Run() {
	log.Printf("Server is running on port%s", port)

	runFunction()

	err := http.ListenAndServe(port, router)
	if err != nil {
		if err == http.ErrServerClosed {
			lineError = "myapp, line 25: Server closed"
			server_errors.ErrorFatal(err, lineError)
		}
		lineError = "myapp, line 25: Error on ListenAndServe"
		server_errors.ErrorFatal(err, lineError)
	}
}

func runFunction() {
	handler.Handler(router)
}
