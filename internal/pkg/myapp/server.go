package myapp

import (
	"RegisterUser/internal/app/database"
	"RegisterUser/internal/app/handler"
	"fmt"
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
	log.Printf("Starting")

	runFunction()

	err := http.ListenAndServe(port, router)
	if err != nil {
		if err == http.ErrServerClosed {
			log.Fatal("Server closed", err)
		}
		log.Fatal("Error on ListenAndServe", err)
	}

	fmt.Println(http.StatusText(http.StatusOK))
}

func runFunction() {
	ErrorDB()
	handler.Handler(router)
}

func ErrorDB() {
	_, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting database: %v", err)
	}

	log.Print("Connected to MongoDB OK")
}
