package client_errors

import (
	"log"
	"net/http"
)

// Errors that are not Fatals for the server, such as a login error
func ErrorPanic(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	write, newErr := w.Write([]byte(err.Error()))
	if newErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(write, newErr)
	}
}
