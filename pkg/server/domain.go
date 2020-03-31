package server

import (
	"net/http"
	"strconv"

	"later/pkg/model"
	"later/pkg/service"

	"later/pkg/request"

	"github.com/gin-gonic/gin"
)

func NewDomainServer(manager service.DomainManager) DomainServer {
	return DomainServer{manager}
}

// DomainServer exposes endpoints for domain related REST requests
type DomainServer struct {
	Manager service.DomainManager
}

// RegisterEndpoints defines handlers for endpoints for the domain service
func (server *DomainServer) RegisterEndpoints(router *gin.Engine) {
	router.POST("/domains/create", server.create)

	router.GET("/domains/by-domain", server.byDomain)
	router.GET("/domains/all", server.all)
}

func (server *DomainServer) create(context *gin.Context) {
	var json request.DomainCreateRequestBody

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domain, err := model.NewDomain(
		json.Domain,
		json.ContentType)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createddomain, err := server.Manager.Create(domain)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error_type": "On Insert", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, createddomain)
}

func (server *DomainServer) byDomain(context *gin.Context) {
	domainQuery := context.Query("domain")

	if domainQuery == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter domain is required"})
		return
	}

	domain, err := server.Manager.ByDomain(domainQuery)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Domain does not exist"})
		return
	}

	context.JSON(http.StatusOK, domain)
}

func (server *DomainServer) all(context *gin.Context) {
	limit := context.Query("limit")

	var err error
	var limitint int

	if limit == "" {
		limitint = 100
	} else {
		limitint, err = strconv.Atoi(limit)
	}

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter limit must be a number"})
		return
	}

	domains, err := server.Manager.All(limitint)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, domains)

}
