package handlers

import (
	"fmt"
	"log"
	"net/http"

    "github.com/gorilla/mux"
)

//Handler Functions
func HandleHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("try to serve a home page"))
    w.WriteHeader(http.StatusOK)
    log.Print("Serving butterNote homepage")
}

func PostNote(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Server got Note: %v\n", vars["note"])
    log.Printf("Server got Note: %v\n", vars["note"])
    w.WriteHeader(http.StatusOK)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Server replied with: na bhai na, abhi database nahi hai")
    log.Printf("Server replied with: na bhai na, abhi database nahi hai")
    w.WriteHeader(http.StatusOK)
}