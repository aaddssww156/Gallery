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

var style models.Style

func GetAllStyles(w http.ResponseWriter, r *http.Request) {
	styles := style.GetAll()
	json.NewEncoder(w).Encode(styles)
}

func GetStyle(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	style.ID = id

	style := style.Get()
	json.NewEncoder(w).Encode(style)
}

func SaveStyle(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert style data"))
	}

	if err := json.Unmarshal(data, style); err != nil {
		log.Fatal(err)
	}

	style.Save()
}

func DeleteStyle(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	style.ID = id

	style.Delete()
}

func UpdateStyle(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update style data"))
	}

	if err := json.Unmarshal(data, style); err != nil {
		log.Fatal(err)
	}

	style.Update()
}
