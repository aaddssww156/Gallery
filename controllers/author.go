package controllers

import (
	"encoding/json"
	"gallery/models"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var author models.Author

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors := author.GetAll()
	json.NewEncoder(w).Encode(authors)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	author := author.Get(id)
	json.NewEncoder(w).Encode(author)
}

func SaveAuthor(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert author data"))
	}

	var authorData models.Author

	if err := json.Unmarshal(data, authorData); err != nil {
		log.Fatal(err)
	}

	author.Save(authorData)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	author.Delete(id)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update author data"))
	}

	var authorData models.Author

	if err := json.Unmarshal(data, authorData); err != nil {
		log.Fatal(err)
	}

	author.Update(authorData)
}
