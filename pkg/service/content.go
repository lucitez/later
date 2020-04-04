package service

import (
	"fmt"
	"later/pkg/model"
	"later/pkg/parse"
	"later/pkg/repository"

	"github.com/google/uuid"
)

// Content ...
type Content struct {
	DomainService Domain
	Repository    repository.Content
	Parser        parse.Content
}

// NewContent for wire generation
func NewContent(
	domainService Domain,
	repo repository.Content,
	parser parse.Content,
) Content {
	return Content{
		domainService,
		repo,
		parser,
	}
}

// CreateFromURL calls parse and creates content from parse results
func (service *Content) CreateFromURL(url string) (*model.Content, error) {
	urlDomain := parse.DomainFromURL(url)

	domain := service.DomainService.ByDomain(urlDomain)

	content := service.Parser.ContentFromURL(url, domain)

	if service.Parser.Err() != nil {
		fmt.Println("asdlfas;ldkfja;")
		return nil, service.Parser.Err()
	}

	if err := service.Repository.Insert(content); err != nil {
		return nil, err
	}

	return &content, nil
}

// ByID returns Content by ID
func (service *Content) ByID(id uuid.UUID) *model.Content {
	return service.Repository.ByID(id)
}

// All returns all content within a limit
func (service *Content) All(limit int) []model.Content {
	return service.Repository.All(limit)
}
