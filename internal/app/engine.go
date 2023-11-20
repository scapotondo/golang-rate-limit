package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
}

func NewEngine() Engine {
	e := gin.Default()
	// TODO - here's a place to add middlewares set
	return e
}
