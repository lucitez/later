package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"later.co/pkg/repository"

	"later.co/pkg/repository/domainrepo"
	"later.co/pkg/repository/friendrepo"
	"later.co/pkg/repository/usercontentrepo"

	"later.co/pkg/server/friendserver"
	"later.co/pkg/server/usercontentserver"
)

func main() {
	router := gin.Default()

	db, err := repository.InitDB()

	if err != nil {
		log.Panic(err)
	}

	domainrepo.DB = db
	usercontentrepo.DB = db
	friendrepo.DB = db

	usercontentserver.RegisterEndpoints(router)
	friendserver.RegisterEndpoints(router)

	router.Run(":8000")
}
