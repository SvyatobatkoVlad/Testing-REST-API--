package main

import (
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("error running hsttp server")
	}
}
