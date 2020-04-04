package server

import (
	"github.com/gin-gonic/gin"
)

// Server defines the interface all servers shoud extend
type Server interface {
	RegisterEndpoints(router *gin.Engine)
}
