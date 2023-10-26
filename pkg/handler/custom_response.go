package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type errorResponse struct {
	Message string `json:"message"`
}
type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	slog.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
