package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) idValidator(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидный id")
	}
	if intId == 0 {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидный id")
	}
}
func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "пустой header авторизации")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидный токен авторизации")
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set("userId", userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get("userId")
	if !ok {
		return 0, errors.New("id не найден")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("id не целое число")
	}

	return idInt, nil
}
