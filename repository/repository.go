package repository

import (
	"context"
	"errors"
	"meli-backend/model"
	"meli-backend/utils"

	"go.mongodb.org/mongo-driver/bson"
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

func (r repository) GetConversation(ctx context.Context, id string) (model.Conversation, error) {
	var conver model.Conversation
	err := r.db.
		Collection("llmconversations").
		FindOne(ctx, bson.M{"id": id}).
		Decode(&conver)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Conversation{}, utils.ErrConversationNotFound
		}
		return model.Conversation{}, err
	}
	return conver, nil
}
