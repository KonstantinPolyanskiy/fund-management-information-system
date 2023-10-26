package handler

import (
	"fund-management-information-system/internal_types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
