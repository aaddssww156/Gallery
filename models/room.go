package models

import (
	"context"
	"gallery/db"
	"log"

	"github.com/jackc/pgx/v5"
)

// Объект для хранения информации о комнатах
// В комнатах есть ограничение количества мест для коллекционеров, которые выставляют в комнатах свои картины
// Комнаты так же могут быть тематическими, например: "Городской пейзаж", "Военная тематика"
type Room struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
	Theme    string `json:"theme,omitempty"`
}

// Удаление выбранной записи tech из таблицы по id
func (r *Room) Delete(id int) error {
	sql := `DELETE FROM room WHERE id=$1`
	_, err := db.DB.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

// Получение выбранной записи tech по id
func (r *Room) Get(id int) Room {
	sql := `SELECT * FROM room WHERE id=$1`

	row, err := db.DB.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	room, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Room])
	if err != nil {
		return Room{}
	}

	return room
}

// Запись сущности tech в базу данных
func (r *Room) Save(room Room) {
	sql := `INSERT INTO room (name, capacity, theme) VALUES ($1, $2, $3)`

	_, err := db.DB.Exec(context.Background(), sql, room.Name, room.Capacity, room.Theme)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func (r *Room) Update(room Room) {
	sql := `UPDATE room SET name = $1, capacity = $2, theme = $3 WHERE id = $4`

	_, err := db.DB.Exec(context.Background(), sql, room.Name, room.Capacity, room.Theme, room.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func (r *Room) GetAll() []Room {
	sql := `SELECT * FROM room`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	room, err := pgx.CollectRows(row, pgx.RowToStructByName[Room])
	if err != nil {
		log.Fatal(err)
	}

	return room
}
