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
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	tech := tech.Get(id)
	json.NewEncoder(w).Encode(tech)
}

func SaveTech(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert tech data"))
	}

	var techData models.Tech

	if err := json.Unmarshal(data, techData); err != nil {
		log.Fatal(err)
	}

	tech.Save(techData)
}

func DeleteTech(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	tech.Delete(id)
}

func UpdateTech(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update tech data"))
	}

	var techData models.Tech

	if err := json.Unmarshal(data, techData); err != nil {
		log.Fatal(err)
	}

	tech.Update(techData)
}
