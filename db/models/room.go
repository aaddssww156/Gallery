package db

import (
	"context"
	"gallery/db"
	"gallery/models"
	"log"

	"github.com/jackc/pgx/v5"
)

// Удаление выбранной записи tech из таблицы по id
func DeleteRoom(id int) error {
	sql := `DELETE FROM room WHERE id=$1`
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
func SaveRoom(name string, capacity int, theme string) {
	sql := `INSERT INTO room (name, capacity, theme) VALUES ($1, $2, 3)`

	_, err := db.DB.Exec(context.Background(), sql, name, capacity, theme)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func UpdateRoom(name, capacity, theme string, id int) {
	sql := `UPDATE room SET name = $1, capacity = $2, theme = $3 WHERE id = $4`

	_, err := db.DB.Exec(context.Background(), sql, name, capacity, theme, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func GetAllRooms() []models.Room {
	sql := `SELECT * FROM room`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	room, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Room])
	if err != nil {
		log.Fatal(err)
	}

	return room
}
