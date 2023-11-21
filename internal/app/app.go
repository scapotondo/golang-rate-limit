package app

import (
	"context"
	"fmt"
	"golang-rate-limit/internal/constants"
	"golang-rate-limit/internal/controller"
	"golang-rate-limit/internal/logs"
	"golang-rate-limit/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	Engine Engine
	Logger logs.Logger

	HealthController       controller.Health
	NotificationService    service.Notification
	NotificationController controller.Notification
}

func (app *App) Setup() *App {
	app.Logger.InfoWithoutContext("App setup started")
	app.injectDependencies()
	app.ConfigureRoutes()
	return app
}

func (app *App) injectDependencies() {
	app.Logger.InfoWithoutContext("Inject dependencies")

	app.HealthController = controller.NewHealth()
	app.NotificationService = service.NewNotification()
	app.NotificationController = controller.NewNotification(app.NotificationService)
}

func (app *App) ConfigureRoutes() {
	app.Logger.InfoWithoutContext("Configuring routes")

	app.Engine = NewEngine()
	v1 := app.Engine.Group("/v1")

	{
		// Swagger
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// Health
		v1.GET(constants.PingBasePath, app.HealthController.GetPing)

		// // Notification
		v1.POST(constants.NotificationBasePath, app.NotificationController.SendEmail)
	}
}

func (app *App) InitServer() {
	app.Logger.InfoWithoutContext("Starting server")

	server := app.createServer()
	go func() {
		app.Logger.InfoWithoutContext("Server started")
		err := server.ListenAndServe()
		if err != nil {
			app.Logger.ErrorWithoutContext("Error starting server", err)
			return
		}
	}()

	app.waitForShutdownSignal(server)
}

func (app *App) createServer() *http.Server {
	serverAddress := fmt.Sprintf(":%d", constants.ServerPort)

	return &http.Server{
		Addr:         serverAddress,
		WriteTimeout: time.Second * 180,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      app.Engine,
	}
}

func (app *App) waitForShutdownSignal(srv *http.Server) {
	shutdownTimeout := time.Duration(100)
	channel := make(chan os.Signal, 1)
	// We will accept grafully shutdowns when we want to quit the app via SIGTERM
	signal.Notify(channel, syscall.SIGTERM)
	// Block until we receive out signal
	<-channel

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	// Doesn't block if there's no connection but will wait otherwise
	// until the timeout deadline
	_ = srv.Shutdown(ctx)
}

func NewApp() *App {
	logger := logs.New("app")

	return &App{
		Logger: logger,
	}
}
