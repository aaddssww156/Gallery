package db

import (
	"context"
	"fmt"
	"gallery/models"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

const DB_URL = "postgres://user:supersecret@localhost:5432/gallery"

var db *pgx.Conn
var err error

func Connect() {
	db, err = pgx.Connect(context.Background(), DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

func DeteleTech(id int) error {
	sql := `DELETE FROM tech WHERE id=$1`
	_, err := db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}
	return nil
}

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
