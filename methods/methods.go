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
	json.NewEncoder(w).Encode(s)
}

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
	json.NewEncoder(w).Encode(s)
}

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
		s   []string
		Day string
	)

	w.WriteHeader(200)
	rows, err := db.Query("SELECT DISTINCT Day FROM posts WHERE Year =? AND Month =? ", params["Year"], params["Month"])
	defer rows.Close()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	for rows.Next() {
		err := rows.Scan(&Day)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		s = append(s, Day)
	}
	json.NewEncoder(w).Encode(s)
}

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
	err := db.QueryRow("SELECT * FROM posts WHERE Year =? AND Month =? AND Day =?", params["Year"], params["Month"], params["Day"]).Scan(&ID, &Body, &Title, &Year, &Month, &Day)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(model.Blog{ID: ID, Title: Title, Body: Body, Year: Year, Month: Month, Day: Day})
}
