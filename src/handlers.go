package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	log.Println("Request path /  httpCode:200")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	// os.Setenv("DB", "admin:password@11.57.69.69")
	// os.Setenv("APIKEY", "P99zE8YiLm")
	todos := Todos{
		Todo{Version: "1.0.0",
			Hostname: getHostname(),
			Time:     time.Now(),
			// Db:       os.Getenv("DB"),
			// Apikey:   os.Getenv("APIKEY"),
			Db:     "admin:password@11.57.69.69",
			Apikey: "P99zE8YiLm",
		},
	}
	//add commanet
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	} else {
		log.Println("Request path /version  httpCode:200")
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	} else {
		return name
	}
	return ""
}
