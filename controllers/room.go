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
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	room := room.Get(id)
	json.NewEncoder(w).Encode(room)
}

func SaveRoom(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't insert room data"))
	}

	var roomData models.Room

	if err := json.Unmarshal(data, &roomData); err != nil {
		log.Fatal(err)
	}

	room.Save(roomData)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	room.Delete(id)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't update room data"))
	}

	var roomData models.Room

	if err := json.Unmarshal(data, &roomData); err != nil {
		log.Fatal(err)
	}

	room.Update(roomData)
}
