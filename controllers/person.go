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

var person models.Person

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	persons := person.GetAll()
	json.NewEncoder(w).Encode(persons)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	person := person.Get(id)
	json.NewEncoder(w).Encode(person)
}

func GetPersonInfo(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	personInfo := person.GetInfo(id)
	json.NewEncoder(w).Encode(personInfo)
}

func SavePaintingToPerson(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't save painting to person"))
	}

	var PtP models.PaintingToPerson

	if err := json.Unmarshal(data, &PtP); err != nil {
		log.Fatal(err)
	}

	person.SavePaintingForPerson(PtP)
}

func SavePersonToRoom(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't save person to room"))
	}

	var PtR models.PersonToRoom

	if err := json.Unmarshal(data, &PtR); err != nil {
		log.Fatal(err)
	}

	person.SavePersonToRoom(PtR)
}

func SavePerson(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert person data"))
	}

	var personData models.Person

	if err := json.Unmarshal(data, &personData); err != nil {
		log.Fatal(err)
	}

	person.Save(personData)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	person.Delete(id)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update person data"))
	}

	var personData models.Person

	if err := json.Unmarshal(data, &personData); err != nil {
		log.Fatal(err)
	}

	person.Update(personData)
}
