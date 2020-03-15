package server

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeserUUID Deserializes a UUID query parameter
func DeserUUID(context *gin.Context, paramName string) (*uuid.UUID, error) {
	idString := context.Query(paramName)

	if idString == "" {
		return nil, errors.New("Parameter " + paramName + " is required")
	}

	id, err := uuid.Parse(idString)

	if err != nil {
		return nil, errors.New("Parameter " + paramName + " must be of type UUID")
	}

	return &id, nil
}
