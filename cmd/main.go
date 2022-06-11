package main

import (
	"context"
	"fmt"
	"log"

	"github.com/me-dolan/test/pkg/database"
)

type User struct {
	JWT     string
	Refresh string
}

func main() {
	mongoCient, err := database.NewClient("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	collection := mongoCient.Database("test").Collection("user")
	test1 := User{"123", "123"}

	insertResult, err := collection.InsertOne(context.TODO(), test1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("добавлен", insertResult.InsertedID)
}
