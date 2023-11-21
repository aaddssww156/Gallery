-- -- DROP TABLE painting_to_person;
-- -- DROP TABLE person_to_room;
-- -- DROP TABLE painting;
-- -- DROP TABLE author;
-- -- DROP TABLE person;
-- -- DROP TABLE room;
-- -- DROP TABLE style;
-- -- DROP TABLE tech;

INSERT INTO author (id, name, surname, date_born, date_died, description) VALUES 
(2, 'Микеланджело', 'Буонарроти', '1475-03-06', '1564-02-18', 'Итальянский художник, скульптор, а также архитектор, поэт и мыслитель. Отдавал предпочтение масштабным проектам, его любимой живописной техникой всю жизнь была фреска. Он всегда стремился достичь гармонии в своих произведениях и выбирал для них исключительно серьезные темы.'),
(3, 'Рафаэль', 'Санти', '1483-04-06', '1520-04-06', 'Итальянский живописец и архитектор. Его работы полны спокойствия, а техника основана на сбалансированности, легкости и свежести письма, единстве композиции. Отличительной чертой, позволившей ему достичь успеха, являлась способность буквально впитывать в себя все самое лучшее, что он находил в работах других великих мастеров.'),
(4, 'Рембрандт', 'ван Рейн', '1606-07-12', '1669-10-04', 'Голландский художник, гравер. Великий мастер светотени, он умел фокусировать внимание именно на тех местах, которые хотел выделить.'),
(5, 'Иван', 'Айвазовский', '1817-07-29', '1900-05-02', 'Русский художник-маринист. Говорил, что сюжет картины слагается у него в памяти, как сюжет стихотворения у поэта. Писал картины «по памяти», а не с натуры, считая, что так они получаются более живыми.'),
(6, 'Клод', 'Моне', '1840-11-14', '1926-12-05', 'Французский художник, пейзажист и один из основателей импрессионизма. Использовал технику рельефного мазка – яркие, чистые и смелые движения кистью.'),
(7, 'Винсент', 'ван Гог', '1853-03-30', '1890-07-29', 'Нидерландский художник-постимпрессионист. Главным в живописи для него был цвет – он смешивал основные цвета с дополнительными, что позволяло ему получать поразительный эффект. Ранние работы насыщены яркими и позитивными цветами, а более поздние – трагическими цветовыми диссонансам, с помощью которых он выражал весь тот негатив, который охватил его разум и сердце.'),
(8, 'Эдвард', 'Мунк', '1863-12-12', '1944-01-23', 'Норвежский живописец и график. Рисовал «еле заметными» мазками, используя контрастные цвета зимней, сумрачной гаммы. Ранние работы характеризуются некой мечтательной неопределенностью, а более поздние обращены к темам одиночества и смерти. Зрелому стилю художника присущ глубокий драматизм.'),
(9, 'Пабло', 'Руис Пикассо', '1881-10-25', '1973-04-08', 'Испанский и французский художник, скульптор, один из создателей кубизма. Его стиль написания картин менялся на протяжении всей его жизни.'),
(10, 'Сальвадор', 'Дали', '1904-05-11', '1989-01-23', 'Испанский живописец, один из самых ярких представителей сюрреализма. Использовал эффект оптической иллюзии. Разработал «параноидно-критический метод» иррационального познания, основанный на критической интерпретации цепочек безумных видений.'),
(11, 'Джон', 'Мартин', '1789-07-19', '1854-02-17', 'Британский художник романтик (романтизм) и гравёр. Он работал в необычной манере меццо-тинто (вид гравюры на металле); прославился изображением сцен катастроф. Его полотна заполнены крошечными фигурками среди грандиозных архитектурных сооружений.'),

-- select * from author;
