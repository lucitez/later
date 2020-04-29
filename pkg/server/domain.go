package server

import (
	"net/http"

	"github.com/lucitez/later/pkg/service"

	"github.com/lucitez/later/pkg/request"

	"github.com/gin-gonic/gin"
)

// NewDomain for wire gen
func NewDomain(
	service service.Domain,
) Domain {
	return Domain{
		service,
	}
}

// Domain exposes endpoints for domain related REST requests
type Domain struct {
	Service service.Domain
}

func (s *Domain) Prefix() string {
	return "/domains"
}

// Routes defines the routes for content API
func (s *Domain) Routes(router *gin.RouterGroup) []gin.IRoutes {
	return []gin.IRoutes{
		router.POST("/create", s.create),

		router.GET("/by-domain", s.byDomain),
		router.GET("/all", s.all),
	}
}

func (s *Domain) create(context *gin.Context) {
	var body request.DomainCreateRequestBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	domain, err := s.Service.Create(body.ToDomainCreateBody())

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, domain)
}

func (s *Domain) byDomain(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "domain", kind: Str, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		domainName := qp["domain"].(*string)

		domain := s.Service.ByDomain(*domainName)

		context.JSON(http.StatusOK, domain)
	}
}

func (s *Domain) all(context *gin.Context) {
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
