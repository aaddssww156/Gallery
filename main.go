package main

import (
	"fmt"
	"gallery/db"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

// 1 - Акварель
// 2 - Гуашь
// 3 - Сангина
// 4 - Уголь
// 5 - Карандаш
// 6 - Масло
// 7 - Эпоксидная смола
// 8 - Текстура
// 9 - Fluid Art
// 10 - Спиртовые чернила

func main() {
	log.Println("server")
	db.Connect()
	// http.HandleFunc("/", handleMain)
	// http.HandleFunc("/click", handleClick)
	http.HandleFunc("/admin", handleAdmin)
	http.HandleFunc("/add-tech", handleAddTech)
	http.HandleFunc("/delete-tech", handleDeleteTech)

	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handleAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tech := db.GetAllTech()
		log.Println(tech)
		t := template.Must(template.ParseFiles("templates/admin.html", "templates/admin-tech.html"))
		t.ExecuteTemplate(w, "admin", tech)
	}
}

func handleAddTech(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		log.Println(r.FormValue("tech"))
		// t := template.Must(template.ParseFiles(getTemplatePath("admin.html"), getTemplatePath("header.html")))
		// t.Execute(w, nil)
	}
}

func handleDeleteTech(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.Write([]byte("Введите валидный id"))
	}

	if err := db.DeteleTech(id); err != nil {
		w.Write([]byte(err.Error()))
	}

	if err == nil {
		fmt.Fprintf(w, "Успешно удалена техника рисования с id %d", id)
	}
}

// func handleMain(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		t := template.Must(template.ParseFiles(getTemplatePath("index.html")))
// 		t.Execute(w, nil)
// 	}
// }

// func handleClick(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		log.Println("clicked")
// 	}
// }

func getTemplatePath(fileName string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(wd, "templates", fileName)
}
