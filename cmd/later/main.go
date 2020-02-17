package main

import (
	"github.com/gin-gonic/gin"
	"later.co/pkg/client/user"
)

func main() {
	router := gin.Default()

	user.RegisterEndpoints(router)

	router.Run(":8000")
}
