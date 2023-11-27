package controllers

import (
	"encoding/json"
	"gallery/models"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var painting models.Painting

func GetAllPaintings(w http.ResponseWriter, r *http.Request) {
	paintings := painting.GetAll()
	json.NewEncoder(w).Encode(paintings)
}

func GetPainting(w http.ResponseWriter, r *http.Request) {
	painting := painting.Get()
	json.NewEncoder(w).Encode(painting)
}

func SavePainting(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert painting data"))
	}

	painting.Save()
}

func DeletePainting(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	painting.Delete()
}

func UpdatePainting(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update painting data"))
	}

	painting.Update()
}
