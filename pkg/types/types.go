package types

type Event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type AllEvents []Event

var Events = AllEvents{
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
