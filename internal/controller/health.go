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

// Send godoc
//
//	@Summary		Healthcheck for microservice
//	@Description	This method is useful to make a healthcheck
//	@Tags			healdh
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/health [get]
func (hh *healthController) GetPing(ctx *gin.Context) {
	response := HealthResponse{
		Status: "pong",
	}

	ctx.JSON(http.StatusOK, response)
}

func NewHealth() Health {
	return new(healthController)
}
