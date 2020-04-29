package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucitez/later/pkg/server"
)

var (
	env = ""
)

func main() {

	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Error loading .env file. %s", err.Error())
	} else {
		fmt.Println(os.Getenv("DB_HOST"))
	}

	env = os.Getenv("env")
	if env == "" {
		env = "local"
	}

	// db, err := util.InitDB()

	// if err != nil {
	// 	log.Panic(err)
	// }

	// authServer := InitializeAuth(db)
	// contentServer := InitializeContent(db)
	// domainServer := InitializeDomain(db)
	// friendServer := InitializeFriend(db)
	// friendRequestServer := InitializeFriendRequest(db)
	// shareServer := InitializeShare(db)
	// userContentServer := InitializeUserContent(db)
	// userServer := InitializeUser(db)
	// chatServer := InitializeChat(db)
	// messageServer := InitializeMessage(db)
	testServer := server.NewTest()

	// engine := InitializeServer(db)

	router := gin.Default()

	test := router.Group(testServer.Prefix())
	{
		testServer.Routes(test)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	router.Run(":" + port)

	// engine.Init(
	// 	[]server.RouteGroup{
	// 		// &contentServer,
	// 		// &domainServer,
	// 		// &friendServer,
	// 		// &friendRequestServer,
	// 		// &shareServer,
	// 		// &userContentServer,
	// 		// &userServer,
	// 		// &chatServer,
	// 		// &messageServer,
	// 	},
	// 	[]server.RouteGroup{
	// 		// &authServer,
	// 	},
	// 	[]server.RouteGroup{
	// 		&testServer,
	// 	},
	// )
}
