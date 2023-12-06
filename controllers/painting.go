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

	painting := painting.Get(id)
	json.NewEncoder(w).Encode(painting)
}

func SavePainting(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert painting data"))
	}

	var paintingData models.Painting

	if err := json.Unmarshal(data, &paintingData); err != nil {
		log.Fatal(err)
	}

	painting.Save(paintingData)
}

func DeletePainting(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	painting.Delete(id)
}

func UpdatePainting(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update painting data"))
	}

	var paintingData models.Painting

	if err := json.Unmarshal(data, &paintingData); err != nil {
		log.Fatal(err)
	}

	painting.Update(paintingData)
}
