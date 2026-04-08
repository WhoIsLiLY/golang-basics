package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintf(w, "User ID: %s", userID)
}

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/users/{id}", userHandler).Methods("GET")
	
	fmt.Println("Router server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}