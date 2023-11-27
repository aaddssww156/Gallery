package controllers

import (
	"encoding/json"
	"gallery/db"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAllPaintings(w http.ResponseWriter, r *http.Request) {
	paintings := db.GetAllPaintings()
	json.NewEncoder(w).Encode(paintings)
}

func GetPainting(w http.ResponseWriter, r *http.Request) {

}

func SavePainting(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert painting data"))
	}

	db.SavePainting(data)
}

func DeletePainting(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	db.DeletePainting(id)
}

func UpdatePainting(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update painting data"))
	}

	db.UpdatePainting(data)
}
