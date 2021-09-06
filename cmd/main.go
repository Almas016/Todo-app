package main

import (
	"log"

	"github.com/Almas016/Todo-app"
	"github.com/Almas016/Todo-app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err!= nil{
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
