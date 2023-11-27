package controllers

import (
	"encoding/json"
	"gallery/db"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	person := db.GetAllTech()
	json.NewEncoder(w).Encode(person)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {

}

func SavePerson(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert person data"))
	}

	db.SavePerson(data)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	db.DeletePerson(id)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update person data"))
	}

	db.UpdatePerson(data)
}
