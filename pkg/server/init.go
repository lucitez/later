package server

import (
	"log"
	"net/http"
	"os"

	"github.com/lucitez/later/pkg/util/env"

	"github.com/lucitez/later/pkg/auth"

	"github.com/gin-gonic/gin"
)

type RouteGroup interface {
	Prefix() string
	Routes(c *gin.RouterGroup) []gin.IRoutes
}

type Server struct {
	AuthService auth.Service
}

func NewServer(authService auth.Service) Server {
	server := Server{
		AuthService: authService,
	}

	return server
}

func (s *Server) Init(
	protectedRoutes []RouteGroup,
	unprotectedRoutes []RouteGroup,
	testRoutes []RouteGroup,
) {
	router := gin.Default()

	for _, routeGroup := range protectedRoutes {
		protected := router.Group(routeGroup.Prefix(), s.protectedAuth())
		{
			routeGroup.Routes(protected)
		}
	}

	for _, routeGroup := range unprotectedRoutes {
		protected := router.Group(routeGroup.Prefix(), unprotected())
		{
			routeGroup.Routes(protected)
		}
	}

	for _, routeGroup := range testRoutes {
		test := router.Group(routeGroup.Prefix())
		{
			routeGroup.Routes(test)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	router.Run(":" + port)
}

func unprotected() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Search user in the slice of allowed credentials
		clientID := c.GetHeader("Client-ID")

		if clientID != env.MustGetenv("CLIENT_ID") {
			c.Header("WWW-Authenticate", "Basic realm=Authorization Required")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

// protectedAuth requires you to have a valid userSession_id in your access_token. Gives user_id of the userSession to the context
func (s *Server) protectedAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := auth.ParseToken(c.GetHeader("Authorization"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		activeUserSession, err := s.AuthService.ActiveByID(token.UserSessionID)

		if activeUserSession == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "userSession expired")
			return
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		if activeUserSession.UserID.String() != token.StandardClaims.Subject {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Session does not match user_id")
			return
		}

		c.Set("user_id", activeUserSession.UserID)
	}
}
