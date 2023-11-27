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

var room models.Room

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms := room.GetAll()
	json.NewEncoder(w).Encode(rooms)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	room.ID = id

	room := room.Get()
	json.NewEncoder(w).Encode(room)
}

func SaveRoom(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert room data"))
	}

	if err := json.Unmarshal(data, room); err != nil {
		log.Fatal(err)
	}

	room.Save()
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	room.ID = id

	room.Delete()
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update room data"))
	}

	if err := json.Unmarshal(data, room); err != nil {
		log.Fatal(err)
	}

	room.Update()
}
