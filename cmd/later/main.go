package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/lucitez/later/pkg/inits"
	"github.com/lucitez/later/pkg/server"
)

func main() {

	var env string
	var err error
	flag.StringVar(&env, "env", "prod", "specify environment")

	flag.Parse()

	switch env {
	case "local":
		err = godotenv.Load(".env.local")
	case "stage":
		err = godotenv.Load(".env.stage")
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := inits.DB()

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
	testServer := server.NewTest()

	engine := InitializeServer(db)

	engine.Init(
		[]server.RouteGroup{
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
		[]server.RouteGroup{
			&authServer,
		},
		[]server.RouteGroup{
			&testServer,
		},
	)
}
