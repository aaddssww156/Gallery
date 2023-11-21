package models

import "time"

// Словарь для хранения техник написания картин
type Tech struct {
	ID   int    `json:"id"`
	Tech string `json:"tech"`
}

// Словарь для хранения стиля написания картин
type Style struct {
	ID    int    `json:"id"`
	Style string `json:"style"`
}

// Объект для хранения информации о комнатах
// В комнатах есть ограничение количества мест для коллекционеров, которые выставляют в комнатах свои картины
// Комнаты так же могут быть тематическими, например: "Городской пейзаж", "Военная тематика"
type Room struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
	Theme    string `json:"theme,omitempty"`
}

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

// Объект для хранения информации о картинах
// Объект связан со стилями и техникой написания, а так же автором, который написал эту картину
// TODO: добавить путь до файла на сервере для отображения изображения на клиенте
type Painting struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Year     int    `json:"year,omitempty"`
	AuthorId int    `json:"author_id,omitempty"`
	StyleId  int    `json:"style_id,omitempty"`
	TechId   int    `json:"tech_id,omitempty"`
}

// Объект для хранения информации о коллекционере.
// Если коллекционер не активен, то он архивируется и его учетная запись не используется на сайте
type Person struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Email  string `json:"email,omitempty"`
	Active bool   `json:"active,omitempty"`
}
