package server

import (
	"errors"
	"later/pkg/util/stringutil"
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
	// Bool bool
	Bool
)

// QueryParameter wraps the name and kind
type QueryParameter struct {
	name     string
	kind     Kind
	fallback *string
	required bool
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
		var value interface{}
		var err error

		switch qp.kind {
		case Str:
			value, err = deserString(d.Context, qp)
		case Int:
			value, err = deserInt(d.Context, qp)
		case Bool:
			value, err = deserBool(d.Context, qp)
		case UUID:
			value, err = deserUUID(d.Context, qp)
		}

		if err != nil {
			d.err = err
			d.valid = false
			break
		} else {
			m[qp.name] = value
		}
	}
	if d.err != nil {
		d.Context.JSON(http.StatusBadRequest, d.err.Error())
	}
	return m, d.valid
}

// DeserUUID Deserializes a UUID query parameter
func deserUUID(
	context *gin.Context,
	param QueryParameter,
) (*uuid.UUID, error) {
	idString := context.Query(param.name)

	switch {
	case idString == "" && param.required:
		return nil, errors.New("Parameter " + param.name + " is required")
	case idString == "" && !param.required:
		return nil, nil
	default:
		id, err := uuid.Parse(idString)

		if err != nil {
			return nil, errors.New("Parameter " + param.name + " must be a UUID")
		}

		return &id, nil
	}
}

// DeserString Deserializes a String query parameter
func deserString(
	context *gin.Context,
	param QueryParameter,
) (*string, error) {
	var str string

	switch {
	case param.fallback != nil:
		str = context.DefaultQuery(param.name, *param.fallback)
	default:
		str = context.Query(param.name)
	}

	if str == "" && param.required {
		return nil, errors.New("Parameter " + param.name + " is required")
	}

	return stringutil.NullIfBlank(str), nil
}

// DeserInt Deserializes a Int query parameter
func deserInt(
	context *gin.Context,
	param QueryParameter,
) (*int, error) {
	var str string

	switch {
	case param.fallback != nil:
		str = context.DefaultQuery(param.name, *param.fallback)
	default:
		str = context.Query(param.name)
	}

	switch {
	case str == "" && param.required:
		return nil, errors.New("Parameter " + param.name + " is required")
	case str == "" && !param.required:
		return nil, nil
	default:
		strInt, err := strconv.Atoi(str)

		if err != nil {
			return nil, errors.New("Parameter " + param.name + " must be an integer")
		}

		return &strInt, nil
	}
}

// DeserBool Deserializes a Int query parameter
func deserBool(
	context *gin.Context,
	param QueryParameter,
) (*bool, error) {
	var str string

	switch {
	case param.fallback != nil:
		str = context.DefaultQuery(param.name, *param.fallback)
	default:
		str = context.Query(param.name)
	}

	switch {
	case str == "" && param.required:
		return nil, errors.New("Parameter " + param.name + " is required")
	case str == "" && !param.required:
		return nil, nil
	default:
		strBool, err := strconv.ParseBool(str)

		if err != nil {
			return nil, errors.New("Parameter " + param.name + " must be a boolean")
		}

		return &strBool, nil
	}
}
