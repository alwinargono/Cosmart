package main

import (
	"fmt"
	"log"
	"net/http"
	"test/controller"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	fmt.Printf("Welcome to The Public Library\n")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getlist", controller.GetList).Methods("POST")
	router.HandleFunc("/pickup", controller.PickUp).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
