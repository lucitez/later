package repository_test

import (
	"later/pkg/model"
	"later/pkg/repository"
	"later/pkg/repository/util"
	"testing"
)

var domainRepo repository.Domain

var domain, _ = model.NewDomain(
	"test_domain",
	"images")

func TestDomainInsertAndByDomain(t *testing.T) {
	beforeEach()
	domainRepo.Insert(domain)

	actual, _ := domainRepo.ByDomain("test_domain")

	util.AssertEquals(t, actual, domain)
}

func TestDomainAll(t *testing.T) {
	beforeEach()
	domainRepo.Insert(domain)

	actual, _ := domainRepo.All(1)

	util.AssertContainsOne(t, actual, *domain)
}
