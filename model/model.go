package model

type Blog struct {
	ID    string `json: "ID"`
	Year  string `json: "Year"`
	Month string `json: "Month"`
	Day   string `json: "Day"`
	Title string `json: "Title"`
	Body  string `json: "Body"`
}

type Comment struct {
	ID    string `json: "ID"`
	Email string `json: "Email"`
	Alias string `json: "Alias"`
	Body  string `json: "Body"`
}

type List struct {
	Data []string `json: "Data"`
}
