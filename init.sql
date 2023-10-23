-- Словарь с возможными стилями живописи
CREATE TABLE IF NOT EXISTS style (
    id SERIAL NOT NULL PRIMARY KEY ,
    style TEXT
);

-- Словарь с возможными стилями рисования
CREATE TABLE IF NOT EXISTS tech (
    id SERIAL NOT NULL PRIMARY KEY,
    tech TEXT
);

-- Комната для проведения выставки
CREATE TABLE IF NOT EXISTS room (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    capacity INT,
    theme VARCHAR(255)
);

-- Художник, написавший картину
CREATE TABLE IF NOT EXISTS author (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    surname VARCHAR(255),
    date_born timestamp without time zone,
    date_died timestamp without time zone,
    description TEXT
);

-- Картина
CREATE TABLE IF NOT EXISTS painting (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    year INT,
    author_id int references author(id),
    style_id int references style(id),
    tech_id int references tech(id)
);

-- Владелец картины
CREATE TABLE IF NOT EXISTS person (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    phone VARCHAR(30),
    email VARCHAR(255),
    active boolean
);

-- M:M владельцев и комнат
CREATE TABLE IF NOT EXISTS person_to_room (
    room_id int references room(id),
    person_id int references person(id)
);

-- M:M картин и владельцев
CREATE TABLE IF NOT EXISTS painting_to_person (
    painting_id int references painting(id),
    person_id int references person(id)
);