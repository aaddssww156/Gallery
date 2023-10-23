package main

import (
	"gallery/db"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

// var t = template.Must(template.ParseGlob("templates/*"))

func main() {
	log.Println("server")
	db.Connect()
	// http.HandleFunc("/", handleMain)
	// http.HandleFunc("/click", handleClick)
	http.HandleFunc("/admin", handleAdmin)
	http.HandleFunc("/add-tech", handleAddTech)

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
