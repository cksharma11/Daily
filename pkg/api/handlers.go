package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cksharma11/daily/pkg/types"
	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task types.Task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter task is correct format")
	}

	json.Unmarshal(reqBody, &task)
	types.Tasks = append(types.Tasks, task)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(task)
}

func GetOneTask(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)["id"]

	for _, task := range types.Tasks {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
		}
	}
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(types.Tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task types.Task
	taskID := mux.Vars(r)["id"]
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter correct task structure")
	}

	json.Unmarshal(reqBody, &task)

	for i, t := range types.Tasks {
		if t.ID == taskID {
			types.Tasks[i].Date = task.Date
			types.Tasks[i].Description = task.Description
			json.NewEncoder(w).Encode(t)
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)["id"]
	for i, task := range types.Tasks {
		if task.ID == taskID {
			types.Tasks = append(types.Tasks[:i], types.Tasks[i+1:]...)
			fmt.Fprintf(w, "task with id %s is deleted sucessfuly", taskID)
		}
	}
}

func MarkTaskDone(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)["id"]
	for i, task := range types.Tasks {
		if task.ID == taskID {
			types.Tasks[i].Done = true
			json.NewEncoder(w).Encode(task)
		}
	}
}

func HelloAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello API")
}
