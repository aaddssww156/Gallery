package main

import (
	"encoding/json"
	"gallery/db"
	"log"
	"net/http"
)

func main() {
	db.Connect()
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/styles/", handleStyle)
	http.HandleFunc("/tech/", handleTech)
	http.HandleFunc("/authors", handleAuthors)

	log.Println("server is running on 3333 port")
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
}

func handleStyle(w http.ResponseWriter, r *http.Request) {
	styles := db.GetAllStyle()
	json.NewEncoder(w).Encode(styles)
}

func handleTech(w http.ResponseWriter, r *http.Request) {
	// log.Println(path.Base(r.URL.String()))
	tech := db.GetAllTech()
	json.NewEncoder(w).Encode(tech)
}

func handleRooms(w http.ResponseWriter, r *http.Request) {
	// rooms := db.GetAllRooms()
}

func handleAuthors(w http.ResponseWriter, r *http.Request) {
	authors := db.GetAllAuthors()
	json.NewEncoder(w).Encode(authors)
}
