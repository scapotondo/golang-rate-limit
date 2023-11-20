package main

import (
	_ "golang-rate-limit/docs"
	"golang-rate-limit/internal/app"
)

// @title           Golang Rate Limit API
// @version         1.0
// @description     Golang Rate Limit API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  sebastian.capotondo@gmail.com

// @host      localhost:8080
// @BasePath  /v1

func main() {
	app.NewApp().
		Setup().
		InitServer()
}
