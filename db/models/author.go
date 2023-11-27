package db

import (
	"context"
	"gallery/db"
	"gallery/models"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

// Удаление выбранной записи tech из таблицы по id
func DeleteAuthor(id int) error {
	sql := `DELETE FROM author WHERE id=$1`
	_, err := db.DB.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

// Получение выбранной записи tech по id
// func GetTech(id int) models.Tech {
// 	sql := `SELECT * FROM tech WHERE id=$1`

// 	row, err := db.DB.Query(context.Background(), sql, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	tech, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Tech])
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return tech
// }

// Запись сущности tech в базу данных
func SaveAuthor(name, surname string, date_born, date_died time.Time, description string) {
	sql := `INSERT INTO author (name, surname, date_born, date_died, description) VALUES ($1, $2, $3, $4, $5)`

	_, err := db.DB.Exec(context.Background(), sql, name, surname, date_born, date_died, description)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func UpdateAuthor(name, surname string, date_born, date_died time.Time, description string, id int) {
	sql := `UPDATE author SET name = $1, surname = $2, date_born = $3, date_died = $4, description = $5 WHERE id = $6`

	_, err := db.DB.Exec(context.Background(), sql, name, surname, date_born, date_died, description, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func GetAllAuthors() []models.Author {
	sql := `SELECT * FROM author`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	authors, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Author])
	if err != nil {
		log.Fatal(err)
	}

	return authors
}
