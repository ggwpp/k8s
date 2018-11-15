package main

import (
	"log"
	"net/http"
)

func main() {

	port := "8080"

	router := NewRouter()
	log.Println("Staring server on port " + port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
