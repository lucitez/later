package repository_test

import (
	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
	"testing"
)

var domainRepo repository.Domain

var domain = model.NewDomain(
	"test_domain",
	"images",
)

func TestDomainInsertAndByDomain(t *testing.T) {
	beforeEach(t)
	domainRepo.Insert(domain)

	actual := domainRepo.ByDomain("test_domain")

	testUtil.Assert.Equal(*actual, domain)
}

func TestDomainAll(t *testing.T) {
	beforeEach(t)
	domainRepo.Insert(domain)

	actual := domainRepo.All(1)

	testUtil.Assert.Contains(actual, domain)
}
