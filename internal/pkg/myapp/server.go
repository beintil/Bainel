package myapp

import (
	"Bainel/internal/app/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

var (
	router = mux.NewRouter()
)

func Run() {
	log.Printf("Server is running on port%s", port)

	runFunction()

	err := http.ListenAndServe(port, router)
	if err != nil {
		if err == http.ErrServerClosed {
			log.Fatal("Server closed", err)
		}
		log.Fatal("Error on ListenAndServe", err)
	}
}

func runFunction() {
	handler.Handler(router)
}
