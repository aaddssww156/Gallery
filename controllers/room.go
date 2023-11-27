package controllers

import (
	"encoding/json"
	"gallery/db"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms := db.GetAllRooms()
	json.NewEncoder(w).Encode(rooms)
}

func SaveRoom(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert room data"))
	}

	db.SaveRoom(data)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	db.DeleteRoom(id)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update room data"))
	}

	db.UpdateRoom(data)
}
