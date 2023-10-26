package handler

import (
	"fund-management-information-system/internal_types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const RequestParameterId = "id"

func (h *Handler) deleteManager(ctx *gin.Context) {
	id, err := getId(ctx, RequestParameterId)
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
	id, err := getId(ctx, RequestParameterId)
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
	var wantManager, updatedManager internal_types.Manager

	id, err := getId(ctx, RequestParameterId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидный id")
		return
	}

	if err = ctx.BindJSON(&wantManager); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "невалидная структура")
		return
	}

	updatedManager, err = h.service.Manager.UpdateManager(id, wantManager)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updatedManager)
}

func (h *Handler) getManagers(ctx *gin.Context) {
	from, err := strconv.Atoi(ctx.Query("from"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "некорректное начальное значение")
		return
	}

	if from <= 0 {
		newErrorResponse(ctx, http.StatusBadRequest, "значение меньше или равно 0")
		return
	}

	managers, err := h.service.Manager.GetManagers(from)

	ctx.JSON(http.StatusOK, managers)
}

func getId(ctx *gin.Context, keyId string) (int, error) {
	id, err := strconv.Atoi(ctx.Param(keyId))
	if err != nil {
		return 0, err
	}
	return id, nil
}
