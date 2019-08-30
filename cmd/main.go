package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joecomscience/alert-manager/pkg"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/line", line).Methods("GET")
	r.HandleFunc("/msteam", msTeam).Methods("GET")
	r.HandleFunc("/sms", sms).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func line(w http.ResponseWriter, r *http.Request) {
	line := pkg.Line{
		Message: "--test--",
	}
	line.Send()
	w.Write([]byte("line"))
}

func msTeam(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ms team"))
}

func sms(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sms"))
}
