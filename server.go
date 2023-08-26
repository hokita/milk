package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	application "github.com/hokita/milk/application/controller"
	infra "github.com/hokita/milk/infra/repository"
)

type Server struct{}

func (s *Server) Start() error {
	dbUser := os.Getenv("DB_USER")
	dbTable := os.Getenv("DB_TABLE")
	dbIP := os.Getenv("DB_IP")
	dbPW := os.Getenv("DB_PASSWORD")

	sourceName := dbUser + ":" + dbPW + "@tcp(" + dbIP + ")/" + dbTable + "?parseTime=true&loc=Asia%2FTokyo"

	// db
	db, err := sql.Open("mysql", sourceName)
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

	// repositories
	helloRepository := &infra.HelloRepository{}
	logRepository := &infra.LogRepository{DB: db}

	// handlers
	http.Handle("/hello/", &application.HelloHandler{
		Repository: helloRepository,
	})
	http.Handle("/logs/", &application.LogHandler{
		Repository: logRepository,
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}
