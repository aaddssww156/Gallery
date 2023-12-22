package controllers

import (
	"context"
	"gallery/db"
	"net/http"
)

func CleanUp(w http.ResponseWriter, r *http.Request) {
	if (!r.URL.Query().Has("pass")) || (r.URL.Query().Get("pass") != "112233") {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		sql := `
		TRUNCATE TABLE style RESTART IDENTITY CASCADE;
		TRUNCATE TABLE tech RESTART IDENTITY CASCADE;
		TRUNCATE TABLE room RESTART IDENTITY CASCADE;
		TRUNCATE TABLE author RESTART IDENTITY CASCADE;
		TRUNCATE TABLE painting RESTART IDENTITY CASCADE;
		TRUNCATE TABLE person RESTART IDENTITY CASCADE;
		TRUNCATE TABLE person_to_room RESTART IDENTITY CASCADE;
		TRUNCATE TABLE painting_to_person RESTART IDENTITY CASCADE;

		`
		_, err := db.DB.Exec(context.Background(), sql)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error cleaning up all data"))
		}
	}
}
