package auth

import (
	"errors"
	"fmt"
	"github.com/lucitez/later/pkg/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenFromUserSession(userSession model.UserSession) (signedAt string, signedRt string, err error) {
	accessToken := Token{
		UserSessionID: userSession.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: userSession.ExpiresAt.Unix(),
			IssuedAt:  userSession.CreatedAt.Unix(),
			Subject:   userSession.UserID.String(),
		},
	}

	refreshToken := Token{
		UserSessionID: userSession.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: userSession.CreatedAt.Add(time.Hour * 24 * 7).Unix(),
			IssuedAt:  userSession.CreatedAt.Unix(),
		},
	}

	signedAt, err = jwt.NewWithClaims(jwt.SigningMethodHS256, &accessToken).SignedString([]byte("secret"))

	if err != nil {
		return
	}

	signedRt, err = jwt.NewWithClaims(jwt.SigningMethodHS256, &refreshToken).SignedString([]byte("secret"))

	if err != nil {
		return
	}

	return
}

// ParseToken extracts an auth.Token from the header.
func ParseToken(authHeader string) (*Token, error) {
	if authHeader == "" {
		return nil, errors.New("Authorization Required")
	}

	authParts := strings.Split(authHeader, " ")

	if len(authParts) != 2 {
		return nil, errors.New("Malformed JWT")
	}

	jwt, err := jwt.ParseWithClaims(authParts[1], &Token{}, KeyFunc)

	// jwt.ParseWithClaims performs validation against the Token such as expiration
	if err != nil {
		return nil, fmt.Errorf("Error parsing jwt: %s", err.Error())
	}

	return jwt.Claims.(*Token), nil
}
