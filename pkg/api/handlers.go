package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cksharma11/daily/pkg/types"
	"github.com/gorilla/mux"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event types.Event
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter the correct data")
	}

	json.Unmarshal(reqBody, &event)
	types.Events = append(types.Events, event)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(event)
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range types.Events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func GetAllEvent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(types.Events)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event types.Event
	eventID := mux.Vars(r)["id"]
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter correct event structure")
	}

	json.Unmarshal(reqBody, &event)

	for i, singleEvent := range types.Events {
		if singleEvent.ID == eventID {
			singleEvent.Title = event.Title
			singleEvent.Description = event.Description
			types.Events = append(types.Events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	for i, event := range types.Events {
		if event.ID == eventID {
			types.Events = append(types.Events[:i], types.Events[i+1:]...)
			fmt.Fprintf(w, "event with id %s is deleted sucessfuly", eventID)
		}
	}
}

func HelloAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello API")
}
