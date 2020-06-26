package methods

import (
	"encoding/json"
	"net/http"
	"webapi/model"
	"webapi/secrets"

	"github.com/gorilla/mux"
)

var (
	db, err = secrets.CreateDB()
)

/* TODO */
func GetComment(w http.ResponseWriter, r *http.Request) {

}

func PostComment(w http.ResponseWriter, r *http.Request) {

}

func GetYears(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = db.Ping()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var (
		s    []string
		Year string
	)

	w.WriteHeader(200)
	rows, err := db.Query("SELECT DISTINCT Year FROM posts")
	defer rows.Close()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	for rows.Next() {
		err := rows.Scan(&Year)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		s = append(s, Year)
	}
	json.NewEncoder(w).Encode(model.List{Data: s})
}

// Grabs all the months I have posted in
func GetYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	Year := mux.Vars(r)["Year"]

	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = db.Ping()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var (
		s     []string
		Month string
	)

	w.WriteHeader(200)
	rows, err := db.Query("SELECT DISTINCT Month FROM posts WHERE Year =?", Year)
	defer rows.Close()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	for rows.Next() {
		err := rows.Scan(&Month)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		s = append(s, Month)
	}
	json.NewEncoder(w).Encode(model.List{Data: s})
}

// Returns an ID, and a Title for the blog post
func GetYearMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = db.Ping()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var (
		s     []model.Preview
		ID    string
		Title string
	)

	w.WriteHeader(200)
	rows, err := db.Query("SELECT ID, Title FROM posts WHERE Year =? AND Month =? ", params["Year"], params["Month"])
	defer rows.Close()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	for rows.Next() {
		err := rows.Scan(&ID, &Title)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		s = append(s, model.Preview{ID: ID, Title: Title})
	}
	json.NewEncoder(w).Encode(model.PreviewList{Data: s})
}

// Grabs everything you need about the post
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = db.Ping()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var (
		ID    string
		Body  string
		Title string
		Year  string
		Month string
		Day   string
	)

	w.WriteHeader(200)
	err := db.QueryRow("SELECT * FROM posts WHERE ID =?", params["id"]).Scan(&ID, &Body, &Title, &Year, &Month, &Day)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(model.Blog{ID: ID, Title: Title, Body: Body, Year: Year, Month: Month, Day: Day})
}

func GetLatest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = db.Ping()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var (
		ID string
	)

	w.WriteHeader(200)
	err := db.QueryRow("SELECT MAX(ID) FROM posts").Scan(&ID)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(model.Single{ID: ID})
}
