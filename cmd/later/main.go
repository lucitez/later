package main

import (
	"later"
	"later/pkg/repository/util"
	"log"
)

func main() {

	db, err := util.InitDB()

	if err != nil {
		log.Panic(err)
	}

	authServer := InitializeAuth(db)
	contentServer := InitializeContent(db)
	domainServer := InitializeDomain(db)
	friendServer := InitializeFriend(db)
	friendRequestServer := InitializeFriendRequest(db)
	shareServer := InitializeShare(db)
	userContentServer := InitializeUserContent(db)
	userServer := InitializeUser(db)
	chatServer := InitializeChat(db)
	messageServer := InitializeMessage(db)

	server := InitializeServer(db)

	server.Init(
		[]later.RouteGroup{
			&contentServer,
			&domainServer,
			&friendServer,
			&friendRequestServer,
			&shareServer,
			&userContentServer,
			&userServer,
			&chatServer,
			&messageServer,
		},
		[]later.RouteGroup{
			&authServer,
		},
	)
}
