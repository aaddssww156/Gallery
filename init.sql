-- Словарь с возможными стилями живописи
CREATE TABLE IF NOT EXISTS style (
    id SERIAL NOT NULL PRIMARY KEY,
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
    date_born date,
    date_died date,
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

CREATE OR REPLACE FUNCTION capitalizer() RETURNS TRIGGER AS $$
    BEGIN 
        NEW.name := INITCAP(NEW.name);
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

-- Триггер для капитализации всех имен и фамилий
CREATE OR REPLACE TRIGGER capitalizer 
BEFORE INSERT ON person 
FOR EACH ROW EXECUTE FUNCTION capitalizer();

-- Процедура для поиска всех 
CREATE OR REPLACE PROCEDURE get_max_room()
LANGUAGE SQL
AS $$
    SELECT * FROM room WHERE id = (
        SELECT room_id FROM person_to_room WHERE person_id = (
            SELECT id FROM PERSON WHERE id = (
                SELECT id FROM (
                    SELECT person_id AS id, painting_id FROM painting_to_person GROUP BY person_id, painting_id HAVING COUNT(painting_id) = (
                        SELECT MAX(id_counter) FROM (
                            SELECT painting_id, COUNT(painting_id) as id_counter FROM painting_to_person GROUP BY painting_id) as id_counter
                    )
                ) id limit 1
            )
        )
    );
$$

-- CREATE OR REPLACE FUNCTION clean_up()
-- RETURNS void
-- LANGUAGE plpgsql 
-- AS $$
-- BEGIN
--     TRUNCATE TABLE style RESTART IDENTITY CASCADE;
--     TRUNCATE TABLE tech RESTART IDENTITY CASCADE;
--     TRUNCATE TABLE room RESTART IDENTITY CASCADE;
--     TRUNCATE TABLE author RESTART IDENTITY CASCADE;
--     TRUNCATE TABLE painting RESTART IDENTITY CASCADE;
--     TRUNCATE TABLE person RESTART IDENTITY CASCADE;
--     TRUNCATE TABLE person_to_room RESTART IDENTITY CASCADE;
--     TRUNCATE TABLE painting_to_person RESTART IDENTITY CASCADE;
-- end $$;