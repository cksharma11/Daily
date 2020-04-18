package types

type Task struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type AllTasks []Task

var Tasks = AllTasks{
	{
		ID:          "1",
		Date:        "2020-04-18T09:10:31.733Z",
		Description: "Task 1",
		Done:        false,
	},
	{
		ID:          "2",
		Date:        "2020-04-18T09:10:31.733Z",
		Description: "Task 2",
		Done:        false,
	},
}
