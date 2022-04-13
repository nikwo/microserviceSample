package handlers

import (
	"github.com/gin-gonic/gin"
	"microserviceSample/pkg/services"
	"net/http"
)

type Handler struct {
	services services.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	engine := gin.New()

	api := engine.Group("/api/v1")
	{
		api.GET("/greeting", h.Greeting)
	}

	return engine
}

func HandlerInit(service *services.Service) *Handler {
	return &Handler{}
}

func (h *Handler) Greeting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Hello",
	})
}
