package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	config "github.com/me-dolan/test/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var pathConfig = "./config/config.yaml"
var pathConfig string

func init() {
	flag.StringVar(&pathConfig, pathConfig, "config/config.yaml", "path to configuration file")
	flag.Parse()
}

func main() {
	//cfg
	config, err := config.NewConfig(pathConfig)
	if err != nil {
		panic(err)
	}

	// mongo
	mongoCient, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		panic(err)
	}
	err = mongoCient.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	// server
	r := gin.Default()
	fmt.Println("Сервер запущен на порте 8080")
	r.Run(config.Port)
}
