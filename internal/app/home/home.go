package home

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("YOUR HOME PAGE"))
}
