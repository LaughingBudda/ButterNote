package handlers

import (
	"fmt"
	"log"
	"net/http"

    "github.com/gorilla/mux"
)

//Handler Functions
func HandleHome(w http.ResponseWriter, r *http.Request) {
    print("try to serve a home page")
    w.WriteHeader(http.StatusOK)
    log.Print("Serving butterNote homepage")
}

func PostNote(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Server got Note: %v\n", vars["note"])
    log.Printf("Server got Note: %v\n", vars["note"])
}

func GetNote(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Server replied with: na bhai, abhi database nahi hai")
    log.Printf("Server replied with: na bhai, abhi database nahi hai")
}