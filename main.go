package main

import (
	"fmt"
	"log"
	"net/http"
	producthandler "server/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", producthandler.GetAllProductsWithTempStatus).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func PrintMe(name string) (string, error) {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
