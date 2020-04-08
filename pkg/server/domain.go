package server

import (
	"net/http"

	"later/pkg/service"

	"later/pkg/request"

	"github.com/gin-gonic/gin"
)

// NewDomain for wire gen
func NewDomain(manager service.Domain) Domain {
	return Domain{manager}
}

// Domain exposes endpoints for domain related REST requests
type Domain struct {
	Manager service.Domain
}

// RegisterEndpoints defines handlers for endpoints for the domain service
func (server *Domain) RegisterEndpoints(router *gin.Engine) {
	router.POST("/domains/create", server.create)

	router.GET("/domains/by-domain", server.byDomain)
	router.GET("/domains/all", server.all)
}

func (server *Domain) create(context *gin.Context) {
	var body request.DomainCreateRequestBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	domain, err := server.Manager.Create(body.ToDomainCreateBody())

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, domain)
}

func (server *Domain) byDomain(context *gin.Context) {
	deser := NewDeser(
		context,
		QueryParameter{name: "domain", kind: Str, required: true},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		domainName := qp["domain"].(*string)

		domain := server.Manager.ByDomain(*domainName)

		context.JSON(http.StatusOK, domain)
	}
}

func (server *Domain) all(context *gin.Context) {
	defaultLimit := "100"

	deser := NewDeser(
		context,
		QueryParameter{name: "limit", kind: Int, fallback: &defaultLimit},
	)

	if qp, ok := deser.DeserQueryParams(); ok {
		limit := qp["limit"].(*int)
		users := server.Manager.All(*limit)

		context.JSON(http.StatusOK, users)
	}
}
