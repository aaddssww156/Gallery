package models

import (
	"context"
	"gallery/db"
	"log"

	"github.com/jackc/pgx/v5"
)

// Объект для хранения информации о коллекционере.
// Если коллекционер не активен, то он архивируется и его учетная запись не используется на сайте
type Person struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Email  string `json:"email,omitempty"`
	Active bool   `json:"active,omitempty"`
	RoomId int    `json:"room_id,omitempty"`
}

type PersonToRoom struct {
	PersonId int `json:"person_id,omitempty"`
	RoomId   int `json:"room_id,omitempty"`
}

type PaintingToPerson struct {
	PaintingId int `json:"painting_id,omitempty"`
	PersonId   int `json:"person_id,omitempty"`
}

type PersonInfo struct {
	PersonName  string `json:"person_name,omitempty"`
	PersonPhone string `json:"person_phone,omitempty"`
	PersonEmail string `json:"person_email,omitempty"`
	RoomName    string `json:"room_name,omitempty"`
}

// Удаление выбранной записи tech из таблицы по id
func (p *Person) Delete(id int) error {
	sql := `DELETE FROM person WHERE id=$1`
	_, err := db.DB.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}

	return nil
}

// Получение выбранной записи tech по id
func (p *Person) Get(id int) Person {
	sql := `SELECT * FROM person WHERE id=$1`

	row, err := db.DB.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	person, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Person])
	if err != nil {
		return Person{}
	}

	return person
}

// Запись сущности tech в базу данных
func (p *Person) Save(person Person) {
	sql := `INSERT INTO person (name, phone, email, active) VALUES ($1, $2, $3, $4)`

	_, err := db.DB.Exec(context.Background(), sql, person.Name, person.Phone, person.Email, person.Active)
	if err != nil {
		log.Fatal(err)
	}
}

// Обновление данных для записи tech
func (p *Person) Update(person Person) {
	sql := `UPDATE person SET name = $1, phone = $2, email = $3, active = $4 WHERE id = $5`

	_, err := db.DB.Exec(context.Background(), sql, person.Name, person.Phone, person.Email, person.Active, person.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех записей tech
func (p *Person) GetAll() []Person {
	sql := `SELECT * FROM room`

	row, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}

	persons, err := pgx.CollectRows(row, pgx.RowToStructByName[Person])
	if err != nil {
		log.Fatal(err)
	}

	return persons
}

func (p *Person) GetInfo(id int) PersonInfo {
	sql := `
		SELECT p.name AS person_name,
			   p.phone AS person_phone,
			   p.email AS person_email,
			   r.name AS room_name
		FROM person p
		JOIN person_to_room ptr ON p.id = ptr.person_id
		JOIN room r ON ptr.room_id = r.id
		WHERE p.id = $1;
	`

	row, err := db.DB.Query(context.Background(), sql, id)
	if err != nil {
		log.Fatal(err)
	}

	var personName string
	var personPhone string
	var personEmail string
	var roomName string
	if row.Next() {
		row.Scan(&personName, &personPhone, &personEmail, &roomName)
	}

	personInfo := PersonInfo{
		PersonName:  personName,
		PersonPhone: personPhone,
		PersonEmail: personEmail,
		RoomName:    roomName,
	}

	return personInfo

}

func (p *Person) SavePersonToRoom(personToRoom PersonToRoom) {
	sql := `INSERT INTO person_to_room (person_id, room_id) VALUES ($1, $2)`

	_, err := db.DB.Exec(context.Background(), sql, personToRoom.PersonId, personToRoom.RoomId)
	if err != nil {
		log.Fatal(err)
	}
}

func (p *Person) SavePaintingForPerson(paintingToPerson PaintingToPerson) {
	sql := `INSERT INTO painting_to_person (painting_id, person_id) VALUES ($1, $2)`

	_, err := db.DB.Exec(context.Background(), sql, paintingToPerson.PaintingId, paintingToPerson.PersonId)
	if err != nil {
		log.Fatal(err)
	}
}
