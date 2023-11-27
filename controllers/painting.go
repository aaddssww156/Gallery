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

var painting models.Painting

func GetAllPaintings(w http.ResponseWriter, r *http.Request) {
	paintings := painting.GetAll()
	json.NewEncoder(w).Encode(paintings)
}

func GetPainting(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	painting.ID = id

	painting := painting.Get()
	json.NewEncoder(w).Encode(painting)
}

func SavePainting(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert painting data"))
	}

	if err := json.Unmarshal(data, painting); err != nil {
		log.Fatal(err)
	}

	painting.Save()
}

func DeletePainting(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	painting.ID = id

	painting.Delete()
}

func UpdatePainting(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update painting data"))
	}

	if err := json.Unmarshal(data, painting); err != nil {
		log.Fatal(err)
	}

	painting.Update()
}
