//+build wireinject

package main

import (
	"database/sql"

	"later.co/pkg/repository"

	"github.com/google/wire"
	"later.co/pkg/manager"
	"later.co/pkg/server"
)

func InitializeDomain(db *sql.DB) server.DomainServer {
	wire.Build(server.NewDomainServer, manager.NewDomainManager, repository.NewDomainRepository)
	return server.DomainServer{}
}
