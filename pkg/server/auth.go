package server

import (
	"github.com/lucitez/later/pkg/auth"
	"github.com/lucitez/later/pkg/request"
	"github.com/lucitez/later/pkg/service"
	"github.com/lucitez/later/pkg/util/stringutil"
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
		router.POST("/logout", server.logout),
		router.POST("/sms-confirmation", server.smsConfirmation),
		router.POST("/sign-up", server.signUp),
		router.POST("/refresh", server.refreshToken),
		router.GET("/sign-up/check-conflicts", server.checkConflicts),
	}
}

// TODO should we check whether there is an active userSession and just return that userSession id?
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

	if accessToken, refreshToken, err := server.startUserSession(user.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func (server *Auth) logout(c *gin.Context) {
	// TODO
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

	if accessToken, refreshToken, err := server.startUserSession(user.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to start UserSession", "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func (server *Auth) smsConfirmation(c *gin.Context) {
	var body request.SMSConfirmationRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	confirmationCode := stringutil.RandomNInt(6)
	content := "Your Later confirmation code is: " + confirmationCode

	service.SendSMS(body.PhoneNumber, content)

	c.JSON(http.StatusOK, confirmationCode)
}

func (server *Auth) refreshToken(c *gin.Context) {
	token, err := auth.ParseToken(c.GetHeader("Authorization"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	// Expire old userSession once we have issued the new one
	// SHOULD WE EVEN DO THIS
	defer server.AuthService.ExpireUserSession(token.UserSessionID)

	userSession, err := server.AuthService.ByID(token.UserSessionID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if userSession == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Could not find userSession"})
		return
	}

	if accessToken, refreshToken, err := server.startUserSession(userSession.UserID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to start UserSession", "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func (server *Auth) checkConflicts(c *gin.Context) {
	deser := NewDeser(
		c,
		QueryParameter{name: "phone_number", kind: Str, required: true},
		QueryParameter{name: "username", kind: Str, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		phoneNumber := qp["phone_number"].(*string)
		username := qp["username"].(*string)

		if err := server.AuthService.CheckConflicts(
			*phoneNumber,
			*username,
		); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.AbortWithStatus(http.StatusOK)
	}
}

/**
 * Create userSession and generate token with userSession id
 */
func (server *Auth) startUserSession(userID uuid.UUID) (signedAt string, signedRt string, err error) {
	userSession, err := server.AuthService.CreateUserSession(userID)

	if err != nil {
		return
	}

	return auth.GenerateTokenFromUserSession(userSession)
}
