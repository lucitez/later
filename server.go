package later

import (
	"later/pkg/auth"
	"net/http"

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

	router.Run(":8000")
}

// TODO get from env
func unprotected() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Search user in the slice of allowed credentials
		clientID := c.GetHeader("Client-ID")

		if clientID != "315aac7e-467f-4acd-b325-71c86f491f54" {
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

		if activeUserSession == nil || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "userSession expired", "message": err.Error()})
			return
		}

		c.Set("user_id", activeUserSession.UserID)
	}
}
