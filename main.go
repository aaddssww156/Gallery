package main

import (
	"gallery/db"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	log.Println("server")

	db.Connect()
	log.Println(db.GetAllTech())
	log.Println(db.GetTech(3))

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/click", handleClick)

	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t := template.Must(template.ParseFiles(getTemplatePath("index.html")))
		t.Execute(w, nil)
	}
	w.WriteHeader(http.StatusNotFound)
}

func handleClick(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		log.Println("clicked")
		w.WriteHeader(200)
	}
}

func getTemplatePath(fileName string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(wd, "templates", fileName)
}
