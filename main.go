package main

import (
	"encoding/json"
	"gallery/db"
	"log"
	"net/http"
)

// Инициализация приложения
// Создаем соединение к базе данных, после чего запускаем миграцию
func init() {
	db.Connect()
	db.Migrate()
}

func main() {
	// Обработчики http запросов
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/tech/", handleTech)
	http.HandleFunc("/styles/", handleStyle)
	http.HandleFunc("/rooms/", handleRooms)
	http.HandleFunc("/authors", handleAuthors)
	http.HandleFunc("/paintings", handlePaintings)
	http.HandleFunc("/persons", handlePersons)

	// Запуск севера
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
	rooms := db.GetAllRooms()
	json.NewEncoder(w).Encode(rooms)
}

func handleAuthors(w http.ResponseWriter, r *http.Request) {
	authors := db.GetAllAuthors()
	json.NewEncoder(w).Encode(authors)
}

func handlePaintings(w http.ResponseWriter, r *http.Request) {
	paintings := db.GetAllPaintings()
	json.NewEncoder(w).Encode(paintings)
}

func handlePersons(w http.ResponseWriter, r *http.Request) {
	persons := db.GetAllPersons()
	json.NewEncoder(w).Encode(persons)
}
