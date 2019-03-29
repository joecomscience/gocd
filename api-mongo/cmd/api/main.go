package main

import (
	"api-mongo/pkg/users"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", users.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", users.CreateUserHandler).Methods("POST")

	log.Printf("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
