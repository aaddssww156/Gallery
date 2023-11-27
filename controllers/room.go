package controllers

import (
	"encoding/json"
	"gallery/models"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var room models.Room

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms := room.GetAll()
	json.NewEncoder(w).Encode(rooms)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	room := room.Get()
	json.NewEncoder(w).Encode(room)
}

func SaveRoom(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert room data"))
	}
	room.Save()
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	room.Delete()
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update room data"))
	}

	room.Update()
}
