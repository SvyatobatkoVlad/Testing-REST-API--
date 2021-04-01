package main

import (
	"log"

	todo "github.com/SvyatobatkoVlad/REST-API"
	handler "github.com/SvyatobatkoVlad/REST-API/pkg/handler"
	"github.com/SvyatobatkoVlad/REST-API/pkg/repository"
	"github.com/SvyatobatkoVlad/REST-API/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error running hsttp server")
	}
}
