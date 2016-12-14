package app

import (
	"net/http"
	"fmt"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if name := r.FormValue("name"); name == "" {
		fmt.Fprintf(w, "Hello world!")
	} else {
		fmt.Fprintf(w, "Hello %s!", name)
	}
}
