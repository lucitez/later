package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Test ...
type Test struct{}

// NewTest for wire generation
func NewTest() Test {
	return Test{}
}

func (server *Test) Prefix() string {
	return "/test"
}

// RegisterEndpoints defines handlers for endpoints for the Test service
func (server *Test) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.GET("", server.test),
	}
}

func (server *Test) test(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, I'm being hosted on GCP")
}
