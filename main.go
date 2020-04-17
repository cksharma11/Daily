package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Intro to rest API",
		Description: "Some random description",
	},
	{
		ID:          "2",
		Title:       "Some more intro to API",
		Description: "Some more random description ",
	},
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var event event
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter the correct data")
	}

	json.Unmarshal(reqBody, &event)
	events = append(events, event)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(event)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", hello)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/event/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events", getAllEvent).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
