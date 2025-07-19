package controllers

import (
	"encoding/json"
	"net/http"

	"lunasphere.co.uk/watchlist/api/services"
	"lunasphere.co.uk/watchlist/api/types"
)

const TEST_USER_ID = 1

func TitlesAPI() {
	http.HandleFunc("/titles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fetchTitles(w, r)
		case http.MethodPost:
			postTitle(w, r)
		case http.MethodPatch:
			editTitle(w, r)
		}
	})
}

func fetchTitles(w http.ResponseWriter, _ *http.Request) {
	titles, err := services.ListTitles(TEST_USER_ID)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(titles)
	if err != nil {
		panic(err)
	}
}

func postTitle(w http.ResponseWriter, r *http.Request) {
	title := types.Title{}

	err := json.NewDecoder(r.Body).Decode(&title)
	if err != nil {
		panic(err)
	}

	err = services.SubmitTitle(title, TEST_USER_ID)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func editTitle(w http.ResponseWriter, r *http.Request) {
	title := types.Title{}

	err := json.NewDecoder(r.Body).Decode(&title)
	if err != nil {
		panic(err)
	}

	err = services.EditTitle(title, TEST_USER_ID)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
