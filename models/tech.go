package models

import (
	"context"
	"gallery/db"
	"log"

	"github.com/jackc/pgx/v5"
)

// Словарь для хранения техник написания картин
type Tech struct {
	ID   int    `json:"id"`
	Tech string `json:"tech"`
}

// Удаление выбранной записи tech из таблицы по id
func (t *Tech) Delete(id int) error {
	sql := `DELETE FROM tech WHERE id=$1`
	_, err := db.DB.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

// Получение выбранной записи tech по id
func (t *Tech) Get(id int) Tech {
	sql := `SELECT * FROM tech WHERE id=$1`

	row, err := db.DB.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	tech, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Tech])
	if err != nil {
		log.Fatal(err)
	}

	return tech
}

// Запись сущности tech в базу данных
func (t *Tech) Save(tech Tech) {
	sql := `INSERT INTO tech (tech) VALUES ($1)`

	_, err := db.DB.Exec(context.Background(), sql, tech.Tech)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func (t *Tech) Update(tech Tech) {
	sql := `UPDATE tech SET tech = $1 WHERE id = $2`

	_, err := db.DB.Exec(context.Background(), sql, tech.Tech, tech.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func (t *Tech) GetAll() []Tech {
	sql := `SELECT * FROM tech`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	tech, err := pgx.CollectRows(row, pgx.RowToStructByName[Tech])
	if err != nil {
		log.Fatal(err)
	}

	return tech
}
