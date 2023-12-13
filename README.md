# Gallery

### О чем проект?
Проект Gallery является учебным проектом, который реализует API сервис веб-сайта онлайн галереи. Серверное приложение взаимодействует с базой данных PostgreSQL, которая разворачивается в Docker контейнере. В последствии будет реализовано приложение на клиентской части

### Как запустить?
- Загрузить репозиторий локально
- Установить Go
- Установить Docker
- Поднять контейнер с базой данных (```make create_db```)
- Установить все зависимости приложения ```go mod tidy```
- Запустить приложение командой ```go run main.go``` или ```make run```

### Технологии проекта
- СУБД: PostgreSQL
- ЯП: Go
- Docker
- Роутер: Chi

### TODO:
- [ ] Дописать makefile
- [ ] Заполнить данные в БД
- [x] Создать процедуры
- [x] Создать триггеры
- [ ] Добавить мигратор
- [ ] Дописать README
- [ ] Добавить CI/CD