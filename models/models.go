package models

import "time"

type Tech struct {
	ID   int    `json:"id"`
	Tech string `json:"tech"`
}

type Style struct {
	ID    int    `json:"id"`
	Style string `json:"style"`
}

type Room struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
	Theme    string `json:"theme,omitempty"`
}

type Author struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Surname     string    `json:"surname,omitempty"`
	Date_born   time.Time `json:"date_born,omitempty"`
	Date_died   time.Time `json:"date_died,omitempty"`
	Description string    `json:"description,omitempty"`
}

type Painting struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Year     int    `json:"year,omitempty"`
	AuthorId int    `json:"author_id,omitempty"`
	StyleId  int    `json:"style_id,omitempty"`
	TechId   int    `json:"tech_id,omitempty"`
}

type Person struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Email  string `json:"email,omitempty"`
	Active bool   `json:"active,omitempty"`
}
