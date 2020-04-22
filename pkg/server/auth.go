package server

import (
	"later/pkg/auth"
	"later/pkg/request"
	"later/pkg/service"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// Auth ...
type Auth struct {
	UserService service.User
	AuthService auth.Service
}

// NewAuth for wire generation
func NewAuth(
	userService service.User,
	authService auth.Service,
) Auth {
	return Auth{
		UserService: userService,
		AuthService: authService,
	}
}

func (server *Auth) Prefix() string {
	return "/auth"
}

// RegisterEndpoints defines handlers for endpoints for the Auth service
func (server *Auth) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.POST("/login", server.login),
		router.POST("/sign-up", server.signUp),
		router.POST("/refresh", server.refreshToken),
	}
}

// TODO should we check whether there is an active session and just return that session id?
func (server *Auth) login(c *gin.Context) {
	identifier, password, ok := c.Request.BasicAuth()

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization Required"})
		return
	}

	/** Lookup user by identifier and password */
	user := server.UserService.ByIdentifierAndPassword(identifier, password)

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "incorrect_credentials"})
		return
	}

	if accessToken, refreshToken, err := server.startSession(user.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to start Session", "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func (server *Auth) signUp(c *gin.Context) {
	phoneNumber, password, ok := c.Request.BasicAuth()

	if !ok {
		c.Header("WWW-Authenticate", "Basic realm=Authorization Required")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var body request.UserSignUpRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := server.UserService.SignUp(body.ToUserSignUpBody(phoneNumber, password))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if accessToken, refreshToken, err := server.startSession(user.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to start Session", "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func (server *Auth) refreshToken(c *gin.Context) {
	token, err := auth.ParseToken(c.GetHeader("Authorization"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	// Expire old session once we have issued the new one
	// SHOULD WE EVEN DO THIS
	defer server.AuthService.ExpireSession(token.SessionID)

	session, err := server.AuthService.ByID(token.SessionID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if session == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Could not find session"})
		return
	}

	if accessToken, refreshToken, err := server.startSession(session.UserID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to start Session", "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

/**
 * Create session and generate token with session id
 */
func (server *Auth) startSession(userID uuid.UUID) (signedAt string, signedRt string, err error) {
	session, err := server.AuthService.CreateSession(userID)

	if err != nil {
		return
	}

	return auth.GenerateTokenFromSession(session)
}
