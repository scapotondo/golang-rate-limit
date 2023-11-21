package controller

import (
	"golang-rate-limit/internal/logs"

	"github.com/gin-gonic/gin"
)

type Notification interface {
	Send(c *gin.Context)
}

type notificationController struct {
	logger logs.Logger
}

// Send godoc
// @Summary Sends a notification
// @Description asasdfasdfasdf
// @Tags notification
// @Accept json
// @Produce json
func (nc *notificationController) Send(c *gin.Context) {

}

func NewNotification() Notification {
	logger := logs.New("Notification Controller")

	return &notificationController{
		logger: logger,
	}
}
