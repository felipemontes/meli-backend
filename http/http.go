package http

import (
	"errors"
	"meli-backend/model"
	"meli-backend/repository"
	"meli-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	repository repository.Repository
}

func NewServer(repository repository.Repository) *Server {
	return &Server{repository: repository}
}

func (s Server) SaveConversation(ctx *gin.Context) {
	var reqConversation model.Conversation
	if err := ctx.ShouldBindJSON(&reqConversation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	err := s.repository.SaveConversation(ctx, reqConversation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (s Server) GetConversation(ctx *gin.Context) {
	convID := ctx.Param("id")
	if convID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation id"})
		return
	}
	conver, err := s.repository.GetConversation(ctx, convID)
	if err != nil {
		if errors.Is(err, utils.ErrConversationNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"conversation": conver})
}
