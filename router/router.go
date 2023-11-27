package router

import (
	"net/http"

	"gallery/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	// Инициализация роутера и настройки запросов
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Получение всех объектов
	router.Get("api/v1/tech", controllers.GetAllTechs)
	router.Get("api/v1/style", controllers.GetAllStyles)
	router.Get("api/v1/room", controllers.GetAllRooms)
	router.Get("api/v1/author", controllers.GetAllAuthors)
	router.Get("api/v1/painting", controllers.GetAllPaintings)
	router.Get("api/v1/person", controllers.GetAllPersons)

	// Получение объекта по id
	router.Get("api/v1/tech/{id}", controllers.GetTech)
	router.Get("api/v1/style/{id}", controllers.GetStyle)
	router.Get("api/v1/room/{id}", controllers.GetRoom)
	router.Get("api/v1/author/{id}", controllers.GetAuthor)
	router.Get("api/v1/painting/{id}", controllers.GetPainting)
	router.Get("api/v1/person/{id}", controllers.GetPerson)

	// Запись объектов в базу данных
	router.Post("api/v1/tech", controllers.SaveTech)
	router.Post("api/v1/style", controllers.SaveStyle)
	router.Post("api/v1/room", controllers.SaveRoom)
	router.Post("api/v1/author", controllers.SaveAuthor)
	router.Post("api/v1/painting", controllers.SavePainting)
	router.Post("api/v1/person", controllers.SavePerson)

	// Удаление объектов с базы данных
	router.Delete("api/v1/tech/{id}", controllers.DeleteTech)
	router.Delete("api/v1/style/{id}", controllers.DeleteStyle)
	router.Delete("api/v1/room/{id}", controllers.DeleteRoom)
	router.Delete("api/v1/author/{id}", controllers.DeleteAuthor)
	router.Delete("api/v1/painting/{id}", controllers.DeletePainting)
	router.Delete("api/v1/person/{id}", controllers.DeletePerson)

	// Удаление объектов с базы данных
	router.Patch("api/v1/tech", controllers.UpdateTech)
	router.Patch("api/v1/style", controllers.UpdateStyle)
	router.Patch("api/v1/room", controllers.UpdateRoom)
	router.Patch("api/v1/author", controllers.UpdateAuthor)
	router.Patch("api/v1/painting", controllers.UpdatePainting)
	router.Patch("api/v1/person", controllers.UpdatePerson)

	return router
}
