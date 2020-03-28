package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"later.co/pkg/repository"
)

func main() {
	router := gin.Default()

	db, err := repository.InitDB()

	if err != nil {
		log.Panic(err)
	}

	domainServer := InitializeDomain(db)

	domainServer.RegisterEndpoints(router)

	router.Run(":8000")
}
