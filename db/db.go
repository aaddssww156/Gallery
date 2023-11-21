package db

import (
	"context"
	"fmt"
	"gallery/models"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

// URL для подключения к базе данных.
// TODO: Вынести в env
const DB_URL = "postgres://user:supersecret@localhost:5432/gallery"

// Структура для взаимодействия с базой данных
var db *pgx.Conn
var err error

// Подключение к базе данных. Если подключение не было установлено, то программа аварийно завершит свою работу
func Connect() {
	db, err = pgx.Connect(context.Background(), DB_URL)
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

	_, err = db.Exec(context.Background(), string(migrateFile))
	if err != nil {
		log.Fatal(err)
	}
}

// Удаление выбранной записи tech из таблицы по id
func DeteleTech(id int) error {
	sql := `DELETE FROM tech WHERE id=$1`
	_, err := db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}
	return nil
}

// Получение выбранной записи tech по id
func GetTech(id int) models.Tech {
	sql := `select * from tech where id=$1`

	row, err := db.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	tech, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Tech])
	if err != nil {
		log.Fatal(err)
	}

	return tech
}

// Получение всех записей tech
func GetAllTech() []models.Tech {
	sql := `select * from tech`

	row, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	tech, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Tech])
	if err != nil {
		log.Fatal(err)
	}

	return tech
}

// Получение всех записей style
func GetAllStyle() []models.Style {
	sql := `select * from style`

	row, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	style, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Style])
	if err != nil {
		log.Fatal(err)
	}

	return style
}

// Получение всех записей room
func GetAllRooms() []models.Room {
	sql := `select * from room`

	row, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	rooms, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Room])
	if err != nil {
		log.Fatal(err)
	}

	return rooms
}

// Получение всех записей authors
func GetAllAuthors() []models.Author {
	sql := `select * from author`

	row, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	authors, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Author])
	if err != nil {
		log.Fatal(err)
	}

	return authors
}

// Получение всех записей painting
func GetAllPaintings() []models.Painting {
	sql := `select * from painting`

	row, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	paintings, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Painting])
	if err != nil {
		log.Fatal(err)
	}

	return paintings
}

// Получение всех записей person
func GetAllPersons() []models.Person {
	sql := `select * from person`

	row, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	persons, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Person])
	if err != nil {
		log.Fatal(err)
	}

	return persons
}
