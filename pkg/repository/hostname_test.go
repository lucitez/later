package repository_test

import (
	"testing"

	"github.com/lucitez/later/pkg/model"
	"github.com/lucitez/later/pkg/repository"
)

var hostnameRepo repository.Hostname

var hostname = model.NewHostname(
	"test_hostname",
	"images",
)

func TestHostnameInsertAndByHostname(t *testing.T) {
	beforeEach(t)
	hostnameRepo.Insert(hostname)

	actual := hostnameRepo.ByHostname("test_hostname")

	testUtil.Assert.Equal(*actual, hostname)
}

func TestHostnameAll(t *testing.T) {
	beforeEach(t)
	hostnameRepo.Insert(hostname)

	actual := hostnameRepo.All(1)

	testUtil.Assert.Contains(actual, hostname)
}
