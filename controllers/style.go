package controllers

import (
	"encoding/json"
	"gallery/models"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var style models.Style

func GetAllStyles(w http.ResponseWriter, r *http.Request) {
	styles := style.GetAll()
	json.NewEncoder(w).Encode(styles)
}

func GetStyle(w http.ResponseWriter, r *http.Request) {
	style := style.Get()
}

func SaveStyle(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert style data"))
	}

	style.Save()
}

func DeleteStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// db.DeleteStyle(id)
	style.Delete()
}

func UpdateStyle(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update style data"))
	}

	style.Update()
}
