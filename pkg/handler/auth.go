package handler

import (
	internal_types "fund-management-information-system/internal_types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignUpClient
// @Tags auth
// @Description Создание и регистрация клиента
// @ID create-client
// @Accept json
// @Produce json
// @Param input body internal_types.SignUpClient true "Данные для создания клиента"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up-client [post]
func (h *Handler) signUpClient(ctx *gin.Context) {
	var input internal_types.SignUpClient

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "некорректный ввод")
		return
	}

	id, err := h.service.Authorization.CreateClient(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary SignUpManager
// @Tags auth
// @Description Создание и регистрация менеджера
// @ID create-manager
// @Accept json
// @Produce json
// @Param input body internal_types.SignUpClient true "Данные для создания менеджера"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up-manager [post]
func (h *Handler) signUpManager(ctx *gin.Context) {
	var input internal_types.ManagerAccount

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "некорректный ввод")
		return
	}

	id, err := h.service.Authorization.CreateManagerAccount(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description Получение jwt токена
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body signInInput true "Данные для авторизации"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "некорректные данные авторизации")
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
