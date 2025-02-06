package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"

	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}