package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"later.co/pkg/repository"

	"later.co/pkg/repository/userrepo"

	"later.co/pkg/server/userserver"
)

func main() {
	router := gin.Default()

	db, err := repository.InitDB()

	if err != nil {
		log.Panic(err)
	}

	userrepo.DB = db

	userserver.RegisterEndpoints(router)

	router.Run(":8000")
}
