package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lucitez/later/pkg/inits"
	"github.com/lucitez/later/pkg/server"
)

func main() {

	var err error

	env := os.Getenv("NODE_ENV")

	fmt.Printf("ENVIRONMENT FROM NODE_ENV: %s", env)

	switch env {
	case "production":
		err = godotenv.Load(".env.prod")
	default:
		err = godotenv.Load(".env.prod")
	}

	if err != nil {
		log.Fatalf("Error loading .env file, %v", err)
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
