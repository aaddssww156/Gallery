package controllers

import (
	"encoding/json"
	"gallery/models"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var person models.Person

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	persons := person.GetAll()
	json.NewEncoder(w).Encode(persons)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	person := person.Get()
	json.NewEncoder(w).Encode(person)
}

func SavePerson(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert person data"))
	}

	person.Save()
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	person.Delete()
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update person data"))
	}

	person.Update()
}
