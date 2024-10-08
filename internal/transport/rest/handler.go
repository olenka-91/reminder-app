package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/olenka--91/reminder-app/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	remind := router.Group("/remind", h.userIdentity)
	{
		remind.POST("/", h.createRemind)
		remind.GET("/", h.getAllReminds)
		remind.GET("/:id", h.getRemindByID)
		remind.PUT("/:id", h.updateRemind)
		remind.DELETE("/:id", h.deleteRemind)
	}

	users := router.Group("/user")
	{
		users.POST("/sign-up", h.signUp)
		users.POST("/sign-in", h.signIn)
	}

	return router

}
