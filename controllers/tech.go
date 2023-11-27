package controllers

import (
	"encoding/json"
	"gallery/db"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAllTechs(w http.ResponseWriter, r *http.Request) {
	tech := db.GetAllTech()
	json.NewEncoder(w).Encode(tech)
}

func SaveTech(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert tech data"))
	}

	db.SaveTech(data)
}

func DeleteTech(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	db.DeleteTech(id)
}

func UpdateTech(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update tech data"))
	}

	db.UpdateTech(data)
}
