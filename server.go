package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LaughingBudda/ButterNote/handlers"
	"github.com/gorilla/mux"
)


func startServer() {
	//initialize the server to start accepting requests
	r := mux.NewRouter()
	r.Host("www.butternote.com")

	//all the API end point
	r.HandleFunc("/", handlers.HandleHome)   //have to add .Methods("GET") to make it REST
	r.HandleFunc("/postnote/{note}", handlers.PostNote)
	r.HandleFunc("/getnote", handlers.GetNote)

	//enable the router
	http.Handle("/", r) 

	//Start server
	port := ":6969"
	fmt.Println("\nListening on port " + port)
	log.Print(http.ListenAndServe(port, r))
}


func main() {
	startServer()
}
