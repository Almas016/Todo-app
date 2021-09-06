package main

import (
	"log"

	todo "github.com/Almas016/Todo-app"
	"github.com/Almas016/Todo-app/pkg/handler"
	"github.com/Almas016/Todo-app/pkg/repository"
	"github.com/Almas016/Todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(*repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
