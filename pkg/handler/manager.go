package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) deleteManager(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидный id")
		return
	}
	err = h.service.Manager.DeleteById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) getManagerById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидный id")
		return
	}
	manager, err := h.service.Manager.GetById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, manager)
}

func (h *Handler) updateManager(ctx *gin.Context) {

}

func getId(ctx *gin.Context, keyId string) (int, error) {
	id, err := strconv.Atoi(ctx.Param(keyId))
	if err != nil {
		return 0, err
	}
	return id, nil
}
