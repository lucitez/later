package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"later.co/pkg/repository/util"
)

func main() {
	router := gin.Default()

	db, err := util.InitDB()

	if err != nil {
		log.Panic(err)
	}

	contentServer := InitializeContent(db)
	domainServer := InitializeDomain(db)
	friendServer := InitializeFriend(db)
	friendRequestServer := InitializeFriendRequest(db)
	shareServer := InitializeShare(db)
	userContentServer := InitializeUserContent(db)
	userServer := InitializeUser(db)

	contentServer.RegisterEndpoints(router)
	domainServer.RegisterEndpoints(router)
	friendServer.RegisterEndpoints(router)
	friendRequestServer.RegisterEndpoints(router)
	shareServer.RegisterEndpoints(router)
	userContentServer.RegisterEndpoints(router)
	userServer.RegisterEndpoints(router)

	router.Run(":8000")
}
