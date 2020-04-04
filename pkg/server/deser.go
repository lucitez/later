package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Kind defines the type of the query param
type Kind int

const (
	// Str string
	Str Kind = iota
	// Int int
	Int
	// UUID uuid.UUID
	UUID
)

// QueryParameter wraps the name and kind
type QueryParameter struct {
	name     string
	kind     Kind
	fallback *string
}

// Deser struct does all the work
type Deser struct {
	QueryParameters []QueryParameter
	Context         *gin.Context
	err             error
	valid           bool
}

// NewDeser constructor for Deser. Pass the context and vararg QueryParameter
func NewDeser(
	context *gin.Context,
	params ...QueryParameter,
) Deser {
	return Deser{
		QueryParameters: params,
		Context:         context,
	}
}

// DeserQueryParams deserializes query params into a map of param -> value
// Sets context with error if there is one so the server using this can just return
// Use:
// parameters, ok := deser.DeserQueryParams(); ok {
// 		do stuff
// }
func (d *Deser) DeserQueryParams() (map[string]interface{}, bool) {
	m := make(map[string]interface{})
	d.valid = true
	for _, qp := range d.QueryParameters {
		switch qp.kind {
		case Str:
			value, err := deserString(d.Context, qp.name, qp.fallback)
			if err != nil {
				d.err = err
				d.valid = false
				break
			} else {
				m[qp.name] = *value
			}
		case Int:
			value, err := deserInt(d.Context, qp.name, qp.fallback)
			if err != nil {
				d.err = err
				d.valid = false
				break
			} else {
				m[qp.name] = *value
			}
		case UUID:
			value, err := deserUUID(d.Context, qp.name)
			if err != nil {
				d.err = err
				d.valid = false
				break
			} else {
				m[qp.name] = *value
			}
		}
	}
	if d.err != nil {
		d.Context.JSON(http.StatusBadRequest, d.err.Error())
	}
	return m, d.valid
}

func (d *Deser) Err() (error, bool) {
	return d.err, d.valid
}

// DeserUUID Deserializes a UUID query parameter
func deserUUID(context *gin.Context, paramName string) (*uuid.UUID, error) {
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

// DeserString Deserializes a String query parameter
func deserString(
	context *gin.Context,
	paramName string,
	fallback *string,
) (*string, error) {
	var str string

	switch {
	case fallback != nil:
		str = context.DefaultQuery(paramName, *fallback)
	default:
		str = context.Query(paramName)
	}

	if str == "" {
		return nil, errors.New("Parameter " + paramName + " is required")
	}

	return &str, nil
}

// DeserInt Deserializes a Int query parameter
func deserInt(
	context *gin.Context,
	paramName string,
	fallback *string,
) (*int, error) {
	var str string

	switch {
	case fallback != nil:
		str = context.DefaultQuery(paramName, *fallback)
	default:
		str = context.Query(paramName)
	}

	if str == "" {
		return nil, errors.New("Parameter " + paramName + " is required")
	}

	strInt, err := strconv.Atoi(str)

	if err != nil {
		return nil, errors.New("Parameter " + paramName + " must be of type Int")
	}

	return &strInt, nil
}
