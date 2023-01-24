package server_errors

import "log"

// Server error handler (Fatals errors). For example: database error, tcp error etc.
func ErrorFatal(err error, lineError string) {
	log.Fatalf("Place of error: %s, Error: %v", lineError, err)
}
