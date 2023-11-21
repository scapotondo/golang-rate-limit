package controller

import (
	"golang-rate-limit/internal/logs"
	"golang-rate-limit/internal/resources"
	"golang-rate-limit/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Notification interface {
	SendEmail(c *gin.Context)
}

type notificationController struct {
	logger              logs.Logger
	notificationService service.Notification
}

// Send godoc
//
//	@Summary		Sends a notification
//	@Description	This method is useful to send an email
//	@Description	Restrictions:
//	@Description	- Status type: not more than 2 per minute for each recipient
//	@Description	- News type: not more than 1 per day for each recipient
//	@Description	- Marketing type: not more than 1 per day for each recipient
//	@Tags			notification
//	@Accept			json
//	@Produce		json
//	@Param			request	query	resources.NotificationRequest	true	"Notification details"
//	@Success		200
//	@Failure		400	"Bad Request"
//	@Failure		429	"Too many requests"
//	@Router			/notification/:type [post]
func (nc *notificationController) SendEmail(ctx *gin.Context) {
	var request resources.NotificationRequest
	if err := ctx.ShouldBind(&request); err != nil {
		nc.logger.Error(ctx, "error in NotificationController#SendEmail: bad request", err)
		ctx.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	param, _ := ctx.Params.Get("type")
	request.Type = param

	err := nc.notificationService.SendEmail(ctx, request)
	if err != nil {
		nc.logger.Error(ctx, "error in NotificationController#SendEmail: too many requests", err)
		ctx.Status(http.StatusTooManyRequests)
		return
	}

	ctx.Status(http.StatusOK)
}

func NewNotification(notificationService service.Notification) Notification {
	logger := logs.New("Notification Controller")

	return &notificationController{
		logger:              logger,
		notificationService: notificationService,
	}
}
