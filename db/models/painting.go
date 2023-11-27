package db

import (
	"context"
	"gallery/db"
	"gallery/models"
	"log"

	"github.com/jackc/pgx/v5"
)

// Удаление выбранной записи tech из таблицы по id
func DeletePainting(id int) error {
	sql := `DELETE FROM painting WHERE id=$1`
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
func SavePainting(name string, year, author_id, style_id, tech_id int) {
	sql := `INSERT INTO painting (name, year, author_id, style_id, tech_id) VALUES ($1, $2, $3, $4)`

	_, err := db.DB.Exec(context.Background(), sql, name, year, author_id, style_id, tech_id)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func UpdatePainting(name string, year, author_id, style_id, tech_id, id int) {
	sql := `UPDATE person SET name = $1, phone = $2, email = $3, active = $4 WHERE id = $5`

	_, err := db.DB.Exec(context.Background(), sql, name, year, author_id, style_id, tech_id, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func GetAllPaintings() []models.Painting {
	sql := `SELECT * FROM painting`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	paintings, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Painting])
	if err != nil {
		log.Fatal(err)
	}

	return paintings
}
