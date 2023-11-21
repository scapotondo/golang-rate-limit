package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health interface {
	GetPing(ctx *gin.Context)
}

type healthController struct{}

type HealthResponse struct {
	Status string `json:"status"`
}

func (hh *healthController) GetPing(ctx *gin.Context) {
	response := HealthResponse{
		Status: "pong",
	}

	ctx.JSON(http.StatusOK, response)
}

func NewHealth() Health {
	return new(healthController)
}
