package main

import (
	"gallery/db"
	"gallery/router"
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
	// // Запуск севера
	log.Println("server is running on 3333 port")
	log.Fatal(http.ListenAndServe(":3333", router.Routes()))
}
