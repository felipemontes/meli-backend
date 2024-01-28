package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"meli-backend/http"
	"meli-backend/repository"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(client.Database("llmconversations"))

	server := http.NewServer(repository)

	router := gin.Default()
	{
		router.POST("/conversation/save", server.SaveConversation)
	}
}
