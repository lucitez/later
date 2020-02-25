package domainserver

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"later.co/pkg/later/domain"
	"later.co/pkg/repository/domainrepo"
	"later.co/pkg/request"
)

// RegisterEndpoints defines handlers for endpoints for the domain service
func RegisterEndpoints(router *gin.Engine) {
	router.POST("/domains/create", create)

	router.GET("/domains/by-domain", byDomain)
	router.GET("/domains/all", all)
}

func create(context *gin.Context) {
	var json request.DomainCreateRequestBody

	err := context.ShouldBindJSON(&json)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domain, err := domain.New(
		json.Domain,
		json.ContentType)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createddomain, err := domainrepo.Insert(domain)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error_type": "On Insert", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, createddomain)
}

func byDomain(context *gin.Context) {
	domainQuery := context.Query("domain")

	if domainQuery == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parameter domain is required"})
		return
	}

	domain, err := domainrepo.ByDomain(domainQuery)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Domain does not exist"})
		return
	}

	context.JSON(http.StatusOK, domain)
}

func all(context *gin.Context) {
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

	domains, err := domainrepo.All(limitint)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, domains)

}
