package repository

import (
	"context"
	"meli-backend/model"
)

type Repository interface {
	SaveConversation(ctx context.Context, newConversation model.Conversation) error
	GetConversation(ctx context.Context, id string) (model.Conversation, error)
	AddMessage(ctx context.Context, convID string, message model.Message) error
}
