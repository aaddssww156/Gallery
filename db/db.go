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

	row := db.QueryRow(context.Background(), sql, id)

	var id2 int
	var tech2 string
	row.Scan(&id2, &tech2)

	return models.Tech{
		ID:   id2,
		Tech: tech2,
	}

}

func GetAllTech() []models.Tech {
	sql := `select * from tech`

	row, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}
	var tech []models.Tech

	for row.Next() {
		var id int
		var techS string
		row.Scan(&id, &techS)
		tech = append(tech, models.Tech{ID: id, Tech: techS})
	}

	return tech
}
func GetAllStyle() []models.Style {
	sql := `select * from style`

	row, err := db.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}
	var style []models.Style

	for row.Next() {
		var id int
		var styleS string
		row.Scan(&id, &styleS)
		style = append(style, models.Style{ID: id, Style: styleS})
	}

	return style
}
