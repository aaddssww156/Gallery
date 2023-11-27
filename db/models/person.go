package db

import (
	"context"
	"gallery/db"
	"gallery/models"
	"log"

	"github.com/jackc/pgx/v5"
)

// Удаление выбранной записи tech из таблицы по id
func DeletePerson(id int) error {
	sql := `DELETE FROM person WHERE id=$1`
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
func SavePerson(name, phone, email string, active bool) {
	sql := `INSERT INTO person (name, phone, email, active) VALUES ($1, $2, $3, $4)`

	_, err := db.DB.Exec(context.Background(), sql, name, phone, email, active)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func UpdatePerson(name, phone, email string, active, bool, id int) {
	sql := `UPDATE person SET name = $1, phone = $2, email = $3, active = $4 WHERE id = $5`

	_, err := db.DB.Exec(context.Background(), sql, name, phone, email, active, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func GetAllPersons() []models.Person {
	sql := `SELECT * FROM room`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	persons, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Person])
	if err != nil {
		log.Fatal(err)
	}

	return persons
}
