package main

import (
	"log"
	"net/http"

	application "github.com/hokita/milk/application/controller"
	infra "github.com/hokita/milk/infra/repository"
)

type Server struct{}

func (s *Server) Start() error {
	helloRepository := &infra.HelloRepository{}
	logRepository := &infra.LogRepository{}

	http.Handle("/hello/", &application.HelloHandler{
		Repository: helloRepository,
	})
	http.Handle("/logs/", &application.LogHandler{
		Repository: logRepository,
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}
