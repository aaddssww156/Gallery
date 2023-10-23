package main

import (
	"context"
	"fmt"
	"gallery/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/jackc/pgx/v5"
)

const DB_URL = "postgres://user:supersecret@localhost:5432/gallery"

func main() {
	log.Println("server")

	conn, err := pgx.Connect(context.Background(), DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var tech models.Tech
	err = conn.QueryRow(context.Background(), "select id, tech from tech where id=$1", 2).Scan(&tech.ID, &tech.Tech)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(tech)

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
