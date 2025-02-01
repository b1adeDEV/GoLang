package main

import (
	"log"
	"todo"
)

func main() {
	server := new(todo.Server)
	if err := server.Run("8000"); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}