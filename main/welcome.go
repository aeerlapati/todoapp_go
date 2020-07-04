package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Canary is main function
func Canary(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Inside Canary\n")

	fmt.Fprint(w, "tweet")
}

func main() {
	fmt.Printf("Starting server at port 9199\n")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/canary", Canary)
	log.Fatal(http.ListenAndServe(":9199", router))
}
