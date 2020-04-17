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

func updateEvent(w http.ResponseWriter, r *http.Request) {
	var event event
	eventID := mux.Vars(r)["id"]
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter correct event structure")
	}

	json.Unmarshal(reqBody, &event)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = event.Title
			singleEvent.Description = event.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	for i, event := range events {
		if event.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "event with id %s is deleted sucessfuly", eventID)
		}
	}
}

func helloAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", helloAPI)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/event/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events", getAllEvent).Methods("GET")
	router.HandleFunc("/event/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/event/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
