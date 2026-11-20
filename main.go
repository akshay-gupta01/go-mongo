package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akshay-gupta01/go-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Println("Starting server at :8000...!")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getSession() *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	return client
}
