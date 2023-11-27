package controllers

import (
	"encoding/json"
	"gallery/db"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors := db.GetAllAuthors()
	json.NewEncoder(w).Encode(authors)
}

func SaveAuthor(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert author data"))
	}

	db.SaveAuthor(data)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	db.DeleteAuthor(id)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update author data"))
	}

	db.UpdateAuthor(data)
}
