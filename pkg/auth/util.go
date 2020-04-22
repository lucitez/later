package auth

import (
	"encoding/base64"
	"errors"
	"later/pkg/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenFromSession(session model.Session) (signedAt string, signedRt string, err error) {
	accessToken := Token{
		SessionID: session.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: session.ExpiresAt.Unix(),
			IssuedAt:  session.CreatedAt.Unix(),
		},
	}

	refreshToken := Token{
		SessionID: session.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: session.CreatedAt.Add(time.Hour * 24 * 7).Unix(),
			IssuedAt:  session.CreatedAt.Unix(),
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

	const prefix = "Basic "

	tokenPart, err := base64.StdEncoding.DecodeString(authHeader[len(prefix):])
	if err != nil {
		return nil, errors.New("Malformed JWT")
	}

	jwt, err := jwt.ParseWithClaims(string(tokenPart), &Token{}, KeyFunc)

	// jwt.ParseWithClaims performs validation against the Token such as expiration
	if err != nil {
		return nil, err
	}

	return jwt.Claims.(*Token), nil
}
