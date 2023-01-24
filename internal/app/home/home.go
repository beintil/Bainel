package home

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("YOUR HOME PAGE"))
	if err != nil {
		panic(err)
	}
}
