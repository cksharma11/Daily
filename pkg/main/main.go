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
	router.HandleFunc("/task", handler.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", handler.GetOneTask).Methods("GET")
	router.HandleFunc("/tasks", handler.GetAllTask).Methods("GET")
	router.HandleFunc("/task/{id}", handler.UpdateTask).Methods("PATCH")
	router.HandleFunc("/task/{id}", handler.DeleteTask).Methods("DELETE")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
