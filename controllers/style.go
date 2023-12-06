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
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	style := style.Get(id)
	json.NewEncoder(w).Encode(style)
}

func SaveStyle(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert style data"))
	}

	var styleData models.Style

	if err := json.Unmarshal(data, &styleData); err != nil {
		log.Fatal(err)
	}

	style.Save(styleData)
}

func DeleteStyle(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	style.Delete(id)
}

func UpdateStyle(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update style data"))
	}

	var styleData models.Style

	if err := json.Unmarshal(data, &styleData); err != nil {
		log.Fatal(err)
	}

	style.Update(styleData)
}
