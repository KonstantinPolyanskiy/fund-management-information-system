package handler

import (
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
