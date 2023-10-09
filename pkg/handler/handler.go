package handler

import (
	"fund-management-information-system/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up-client", h.signUpClient)
		auth.POST("sign-up-manager", h.signUpManager)
		auth.POST("sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		managers := api.Group("/managers")
		{
			managers.GET("/:id", h.getById, h.idValidator)
			managers.DELETE("/:id", h.deleteManager, h.idValidator)
		}
	}

	return router
}
