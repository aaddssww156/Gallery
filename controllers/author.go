package controllers

import (
	"encoding/json"
	"gallery/models"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var author models.Author

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors := author.GetAll()
	json.NewEncoder(w).Encode(authors)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	author := author.Get()
	json.NewEncoder(w).Encode(author)
}

func SaveAuthor(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert author data"))
	}
	author.Save()
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	author.Delete()
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update author data"))
	}

	author.Update()
}
