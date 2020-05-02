package server

import (
	"net/http"

	"github.com/lucitez/later/pkg/service"

	"github.com/lucitez/later/pkg/request"

	"github.com/gin-gonic/gin"
)

// NewHostname for wire gen
func NewHostname(
	service service.Hostname,
) Hostname {
	return Hostname{
		service,
	}
}

// Hostname exposes endpoints for hostname related REST requests
type Hostname struct {
	Service service.Hostname
}

func (s *Hostname) Prefix() string {
	return "/hostnames"
}

// Routes defines the routes for content API
func (s *Hostname) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.POST("/create", s.create),

		router.GET("/by-hostname", s.byHostname),
		router.GET("/all", s.all),
	}
}

func (s *Hostname) create(context *gin.Context) {
	var body request.HostnameCreateRequestBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hostname, err := s.Service.Create(body.ToHostnameCreateBody())

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, hostname)
}

func (s *Hostname) byHostname(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "hostname", kind: Str, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		hn := qp["hostname"].(*string)

		hostname := s.Service.ByHostname(*hn)

		context.JSON(http.StatusOK, hostname)
	}
}

func (s *Hostname) all(context *gin.Context) {
	defaultLimit := "100"

	deser := NewDeser(
		context,
		QueryParameter{name: "limit", kind: Int, fallback: &defaultLimit},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		limit := qp["limit"].(*int)
		users := s.Service.All(*limit)

		context.JSON(http.StatusOK, users)
	}
}
