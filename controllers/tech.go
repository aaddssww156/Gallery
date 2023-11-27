package controllers

import (
	"encoding/json"
	"gallery/models"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var tech models.Tech

func GetAllTechs(w http.ResponseWriter, r *http.Request) {
	techs := tech.GetAll()
	json.NewEncoder(w).Encode(techs)
}

func GetTech(w http.ResponseWriter, r *http.Request) {
	tech := tech.Get()
	json.NewEncoder(w).Encode(tech)
}

func SaveTech(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert tech data"))
	}

	tech.Save()
}

func DeleteTech(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	tech.Delete()
}

func UpdateTech(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update tech data"))
	}

	tech.Update()
}
