package handler

import (
	"fund-management-information-system/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
			managers.GET(":id", h.getManagerById, h.idValidator)
			managers.GET("/group", h.getManagers)
			managers.DELETE(":id", h.deleteManager, h.idValidator)
			managers.PUT(":id", h.updateManager, h.idValidator)
		}
		clients := api.Group("/clients", h.idValidator)
		{
			clients.GET("/:id", h.getClientById, h.idValidator)
			clients.GET("/group", h.getClients)
			clients.DELETE(":id", h.deleteClient)
			clients.PUT(":id", h.updateClient, h.idValidator)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
