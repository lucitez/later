package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"later.co/pkg/repository"

	"later.co/pkg/repository/contentrepo"
	"later.co/pkg/repository/domainrepo"
	"later.co/pkg/repository/sharerepo"
	"later.co/pkg/repository/usercontentrepo"
	"later.co/pkg/repository/userrepo"

	"later.co/pkg/server/contentserver"
	"later.co/pkg/server/friendserver"
	"later.co/pkg/server/shareserver"
	"later.co/pkg/server/usercontentserver"
	"later.co/pkg/server/userserver"
)

func main() {
	router := gin.Default()

	db, err := repository.InitDB()

	if err != nil {
		log.Panic(err)
	}

	userrepo.DB = db
	domainrepo.DB = db
	contentrepo.DB = db
	sharerepo.DB = db
	usercontentrepo.DB = db

	userserver.RegisterEndpoints(router)
	contentserver.RegisterEndpoints(router)
	shareserver.RegisterEndpoints(router)
	usercontentserver.RegisterEndpoints(router)
	friendserver.RegisterEndpoints(router)

	router.Run(":8000")
}
