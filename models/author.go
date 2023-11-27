package models

import (
	"context"
	"gallery/db"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

// Объект для хранения информации о художнике
// В таблице хранится основная информация о художнике, который писал те или иные картины
type Author struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Surname     string    `json:"surname,omitempty"`
	Date_born   time.Time `json:"date_born,omitempty"`
	Date_died   time.Time `json:"date_died,omitempty"`
	Description string    `json:"description,omitempty"`
}

// Удаление выбранной записи tech из таблицы по id
func (a *Author) DeleteAuthor(id int) error {
	sql := `DELETE FROM author WHERE id=$1`
	_, err := db.DB.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

// Получение выбранной записи tech по id
func (a *Author) GetAuthor(id int) Author {
	sql := `SELECT * FROM tech WHERE id=$1`

	row, err := db.DB.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	author, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Author])
	if err != nil {
		log.Fatal(err)
	}

	return author
}

// Запись сущности tech в базу данных
func (a *Author) SaveAuthor(name, surname string, date_born, date_died time.Time, description string) {
	sql := `INSERT INTO author (name, surname, date_born, date_died, description) VALUES ($1, $2, $3, $4, $5)`

	_, err := db.DB.Exec(context.Background(), sql, name, surname, date_born, date_died, description)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func (a *Author) UpdateAuthor(name, surname string, date_born, date_died time.Time, description string, id int) {
	sql := `UPDATE author SET name = $1, surname = $2, date_born = $3, date_died = $4, description = $5 WHERE id = $6`

	_, err := db.DB.Exec(context.Background(), sql, name, surname, date_born, date_died, description, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func (a *Author) GetAllAuthors() []Author {
	sql := `SELECT * FROM author`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	authors, err := pgx.CollectRows(row, pgx.RowToStructByName[Author])
	if err != nil {
		log.Fatal(err)
	}

	return authors
}
