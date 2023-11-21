package controller_test

import (
	"fmt"
	"golang-rate-limit/internal/app"
	"golang-rate-limit/internal/constants"
	"golang-rate-limit/internal/controller"
	"golang-rate-limit/internal/resources"
	"golang-rate-limit/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/h2non/baloo.v3"
)

func TestSendEmailSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := app.NewApp()
	app.Setup()

	notificationService := service.NewNotification()
	notificationController := controller.NewNotification(notificationService)
	app.NotificationController = notificationController

	// Gin attaches dependencies by value so after changing dependencies at app lvl it is required to re-generate gin's engine
	app.ConfigureRoutes()

	testServer := httptest.NewServer(app.Engine)

	request := baloo.New(testServer.URL).
		Post(fmt.Sprintf("/v1/%s/%s", constants.NotificationBasePath, ":type")).
		Params(map[string]string{
			"type": "status",
		}).
		JSON(resources.NotificationRequest{
			Type:    "status",
			User:    "user-1",
			Message: "message",
		}).
		Expect(t)

	_ = request.
		Status(http.StatusOK).
		Done()
}

func TestSendEmailBadRequestError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := app.NewApp()
	app.Setup()

	notificationService := service.NewNotification()
	notificationController := controller.NewNotification(notificationService)
	app.NotificationController = notificationController

	// Gin attaches dependencies by value so after changing dependencies at app lvl it is required to re-generate gin's engine
	app.ConfigureRoutes()

	testServer := httptest.NewServer(app.Engine)

	request := baloo.New(testServer.URL).
		Post(fmt.Sprintf("/v1/%s/%s", constants.NotificationBasePath, ":type")).
		Params(map[string]string{
			"type": "status",
		}).
		JSON(resources.NotificationRequest{
			Type:    "status",
			Message: "message",
		}).
		Expect(t)

	_ = request.
		Status(http.StatusBadRequest).
		Done()
}

func TestSendEmailTooManyRequestsError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := app.NewApp()
	app.Setup()

	notificationService := service.NewNotification()
	notificationController := controller.NewNotification(notificationService)
	app.NotificationController = notificationController

	// Gin attaches dependencies by value so after changing dependencies at app lvl it is required to re-generate gin's engine
	app.ConfigureRoutes()

	testServer := httptest.NewServer(app.Engine)

	for i := 0; i < 3; i++ {
		request := baloo.New(testServer.URL).
			Post(fmt.Sprintf("/v1/%s/%s", constants.NotificationBasePath, ":type")).
			Params(map[string]string{
				"type": "status",
			}).
			JSON(resources.NotificationRequest{
				Type:    "status",
				User:    "user-1",
				Message: "message",
			}).
			Expect(t)

		if i < 2 {
			_ = request.
				Status(http.StatusOK).
				Done()
		} else {
			_ = request.
				Status(http.StatusTooManyRequests).
				Done()
		}
	}
}
