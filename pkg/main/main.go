package main

import (
	"log"
	"net/http"
	"time"

	handler "github.com/cksharma11/daily/pkg/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.HelloAPI)
	router.HandleFunc("/event", handler.CreateEvent).Methods("POST")
	router.HandleFunc("/event/{id}", handler.GetOneEvent).Methods("GET")
	router.HandleFunc("/events", handler.GetAllEvent).Methods("GET")
	router.HandleFunc("/event/{id}", handler.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/event/{id}", handler.DeleteEvent).Methods("DELETE")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
