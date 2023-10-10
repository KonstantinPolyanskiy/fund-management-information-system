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
		managers := api.Group("/managers", h.idValidator)
		{
			managers.GET("/:id", h.getManagerById)
			managers.DELETE("/:id", h.deleteManager)
		}
		clients := api.Group("/clients", h.idValidator)
		{
			clients.DELETE(":id", h.deleteClient)
		}
	}

	return router
}
