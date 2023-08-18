package main

import (
	"log"
	"os"
)

func main() {
	s := Server{}
	if err := s.Start(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
