package db

import (
	"context"
	"gallery/db"
	"gallery/models"
	"log"

	"github.com/jackc/pgx/v5"
)

// Удаление выбранной записи tech из таблицы по id
func DeleteTech(id int) error {
	sql := `DELETE FROM tech WHERE id=$1`
	_, err := db.DB.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

// Получение выбранной записи tech по id
func GetTech(id int) models.Tech {
	sql := `SELECT * FROM tech WHERE id=$1`

	row, err := db.DB.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	tech, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Tech])
	if err != nil {
		log.Fatal(err)
	}

	return tech
}

// Запись сущности tech в базу данных
func SaveTech(tech string) {
	sql := `INSERT INTO tech (tech) VALUES ($1)`

	_, err := db.DB.Exec(context.Background(), sql, tech)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func UpdateTech(tech, id string) {
	sql := `UPDATE tech SET tech = $1 WHERE id = $2`

	_, err := db.DB.Exec(context.Background(), sql, tech, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func GetAllTechs() []models.Tech {
	sql := `SELECT * FROM tech`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	tech, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Tech])
	if err != nil {
		log.Fatal(err)
	}

	return tech
}
