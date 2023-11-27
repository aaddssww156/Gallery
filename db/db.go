package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

// URL для подключения к базе данных.
// TODO: Вынести в env
const DB_URL = "postgres://user:supersecret@localhost:5432/gallery"

// Структура для взаимодействия с базой данных
var DB *pgx.Conn
var err error

// Подключение к базе данных. Если подключение не было установлено, то программа аварийно завершит свою работу
func Connect() {
	DB, err = pgx.Connect(context.Background(), DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

// Функция инициализации \ миграции базы данных. Берет файл init.sql с корня проекта и загружает его в базу данных
// Используется для создания таблиц, связей и изменений структур таблиц
func Migrate() {
	migrateFile, err := os.ReadFile("init.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(context.Background(), string(migrateFile))
	if err != nil {
		log.Fatal(err)
	}
}
