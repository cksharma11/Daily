package types

import "time"

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	Date        time.Time
}

type AllTasks []Task

var Tasks = AllTasks{
	{
		ID:          "1",
		Date:        time.Now(),
		Description: "Task 1",
		Done:        false,
	},
	{
		ID:          "2",
		Date:        time.Now(),
		Description: "Task 2",
		Done:        false,
	},
}
