package repository

import (
	"context"
	"meli-backend/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r repository) SaveConversation(ctx context.Context, newConversation model.Conversation) error {
	_, err := r.db.
		Collection("llmconversations").
		InsertOne(ctx, newConversation)
	if err != nil {
		return err
	}
	return nil
}
