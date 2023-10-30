package handler

import (
	"fund-management-information-system/internal_types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary		DeleteClient
// @Tags			client
// @Description	Удаление клиента
// @Id				delete-client
// @Accept			json
// @Produce		json
// @Param			id		path		int		true	"ID удаляемого клиента"
// @Success		200		{string}	string	"ok"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Security		ApiKeyAuth
// @Router			/api/clients/{id} [delete]
func (h *Handler) deleteClient(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидный id")
		return
	}

	err = h.service.Client.DeleteById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary GetClient
// @Tags client
// @Description получение клиента по его id
// @Id get-client
// @Produce json
// @Param id path int true "ID клиента, которого хочем получить"
// @Success 200 {object} internal_types.Client
// @Failure default {object} errorResponse
// @Security ApiKeyAuth
// @Router /api/clients/{id} [get]
func (h *Handler) getClientById(ctx *gin.Context) {
	id, err := getId(ctx, RequestParameterId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидный id")
		return
	}

	client, err := h.service.Client.GetById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, client)
}
func (h *Handler) updateClient(ctx *gin.Context) {
	var wantClient, updatedClient internal_types.Client

	id, err := getId(ctx, RequestParameterId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидный id")
		return
	}

	if err = ctx.BindJSON(&wantClient); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидная структура клиента")
		return
	}
	updatedClient, err = h.service.Client.UpdateClient(id, wantClient)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updatedClient)
}
func (h *Handler) getClients(ctx *gin.Context) {
	from, err := strconv.Atoi(ctx.Query("from"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "некорректное начальное значение")
		return
	}

	if from <= 0 {
		newErrorResponse(ctx, http.StatusBadRequest, "id меньше или равно 0")
		return
	}

	clients, err := h.service.Client.GetClients(from)

	ctx.JSON(http.StatusOK, clients)
}
