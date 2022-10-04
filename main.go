package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Ping struct {
	Message string `json:"message"`
}

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ping Ping
	if err := json.NewDecoder(r.Body).Decode(&ping); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(ping.Message)

	if err := json.NewEncoder(w).Encode(map[string]string{
		"message": "first",
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	http.ListenAndServe(":9090", mux)

}
