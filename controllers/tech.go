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

var tech models.Tech

func GetAllTechs(w http.ResponseWriter, r *http.Request) {
	techs := tech.GetAll()
	json.NewEncoder(w).Encode(techs)
}

func GetTech(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	tech.ID = id

	tech := tech.Get()
	json.NewEncoder(w).Encode(tech)
}

func SaveTech(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert tech data"))
	}

	if err := json.Unmarshal(data, tech); err != nil {
		log.Fatal(err)
	}

	tech.Save()
}

func DeleteTech(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	tech.ID = id

	tech.Delete()
}

func UpdateTech(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update tech data"))
	}

	if err := json.Unmarshal(data, tech); err != nil {
		log.Fatal(err)
	}

	tech.Update()
}
