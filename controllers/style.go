package controllers

import (
	"encoding/json"
	"gallery/db"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAllStyles(w http.ResponseWriter, r *http.Request) {
	style := db.GetAllStyles()
	json.NewEncoder(w).Encode(style)
}

func SaveStyle(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert style data"))
	}

	db.SaveStyle(data)
}

func DeleteStyle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	db.DeleteStyle(id)
}

func UpdateStyle(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update style data"))
	}

	db.UpdateStyle(data)
}
