package models

import (
	"context"
	"gallery/db"
	"log"

	"github.com/jackc/pgx/v5"
)

// Словарь для хранения стиля написания картин
type Style struct {
	ID    int    `json:"id"`
	Style string `json:"style"`
}

// Удаление выбранной записи tech из таблицы по id
func (s *Style) Delete(id int) error {
	sql := `DELETE FROM style WHERE id=$1`
	_, err := db.DB.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Style) Get(id int) Style {
	sql := `SELECT * FROM style WHERE id=$1`

	row, err := db.DB.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	style, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Style])
	if err != nil {
		log.Fatal(err)
	}

	return style
}

// Запись сущности tech в базу данных
func (s *Style) Save(style string) {
	sql := `INSERT INTO style (style) VALUES ($1)`

	_, err := db.DB.Exec(context.Background(), sql, style)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func (s *Style) Update(style, id string) {
	sql := `UPDATE style SET style = $1 WHERE id = $2`

	_, err := db.DB.Exec(context.Background(), sql, style, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func (s *Style) GetAll() []Style {
	sql := `SELECT * FROM style`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	style, err := pgx.CollectRows(row, pgx.RowToStructByName[Style])
	if err != nil {
		log.Fatal(err)
	}

	return style
}
