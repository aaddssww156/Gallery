package models

import (
	"context"
	"gallery/db"
	"log"

	"github.com/jackc/pgx/v5"
)

// Объект для хранения информации о картинах
// Объект связан со стилями и техникой написания, а так же автором, который написал эту картину
// TODO: добавить путь до файла на сервере для отображения изображения на клиенте
type Painting struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Year     int    `json:"year,omitempty"`
	AuthorId int    `json:"author_id,omitempty"`
	StyleId  int    `json:"style_id,omitempty"`
	TechId   int    `json:"tech_id,omitempty"`
}

// Удаление выбранной записи tech из таблицы по id
func (p *Painting) Delete(id int) error {
	sql := `DELETE FROM painting WHERE id=$1`
	_, err := db.DB.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

// Получение выбранной записи tech по id
func (p *Painting) Get(id int) Painting {
	sql := `SELECT * FROM painting WHERE id=$1`

	row, err := db.DB.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	painting, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Painting])
	if err != nil {
		log.Fatal(err)
	}

	return painting
}

// Запись сущности tech в базу данных
func (p *Painting) Save(painting Painting) {
	sql := `INSERT INTO painting (name, year, author_id, style_id, tech_id) VALUES ($1, $2, $3, $4)`

	_, err := db.DB.Exec(context.Background(), sql, painting.Name, painting.Year, painting.AuthorId, painting.StyleId, painting.TechId)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func (p *Painting) Update(painting Painting) {
	sql := `UPDATE person SET name = $1, phone = $2, email = $3, active = $4 WHERE id = $5`

	_, err := db.DB.Exec(context.Background(), sql, painting.Name, painting.Year, painting.AuthorId, painting.StyleId, painting.TechId, painting.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func (p *Painting) GetAll() []Painting {
	sql := `SELECT * FROM painting`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	paintings, err := pgx.CollectRows(row, pgx.RowToStructByName[Painting])
	if err != nil {
		log.Fatal(err)
	}

	return paintings
}
