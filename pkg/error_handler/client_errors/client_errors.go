package client_errors

import (
	"log"
	"net/http"
)

// Errors that are not Fatals for the server, such as a login error
func ErrorPanic(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, newErr := w.Write([]byte(err.Error()))
	if newErr != nil {
		log.Println("Error writing error message to response:", newErr)
	}
}
