//+build wireinject

package main

import (
	"database/sql"

	"later.co/pkg/parse"
	"later.co/pkg/repository"

	"github.com/google/wire"
	"later.co/pkg/manager"
	"later.co/pkg/server"
)

func InitializeContent(db *sql.DB) server.ContentServer {
	wire.Build(
		server.NewContentServer,
		manager.NewContentManager,
		manager.NewDomainManager,
		repository.NewContentRepository,
		repository.NewDomainRepository,
		parse.NewParser)
	return server.ContentServer{}
}

func InitializeDomain(db *sql.DB) server.DomainServer {
	wire.Build(
		server.NewDomainServer,
		manager.NewDomainManager,
		repository.NewDomainRepository)
	return server.DomainServer{}
}

func InitializeFriend(db *sql.DB) server.FriendServer {
	wire.Build(
		server.NewFriendServer,
		manager.NewFriendManager,
		manager.NewUserManager,
		repository.NewFriendRepository,
		repository.NewUserRepository)
	return server.FriendServer{}
}

func InitializeFriendRequest(db *sql.DB) server.FriendRequestServer {
	wire.Build(
		server.NewFriendRequestServer,
		manager.NewFriendRequestManager,
		repository.NewFriendRequestRepository)
	return server.FriendRequestServer{}
}

func InitializeShare(db *sql.DB) server.ShareServer {
	wire.Build(
		server.NewShareServer,
		manager.NewContentManager,
		manager.NewDomainManager,
		manager.NewShareManager,
		manager.NewUserContentManager,
		manager.NewUserManager,
		repository.NewContentRepository,
		repository.NewDomainRepository,
		repository.NewShareRepository,
		repository.NewUserContentRepository,
		repository.NewUserRepository,
		parse.NewParser)
	return server.ShareServer{}
}

func InitializeUserContent(db *sql.DB) server.UserContentServer {
	wire.Build(
		server.NewUserContentServer,
		manager.NewUserContentManager,
		repository.NewUserContentRepository)
	return server.UserContentServer{}
}

func InitializeUser(db *sql.DB) server.UserServer {
	wire.Build(
		server.NewUserServer,
		manager.NewUserManager,
		repository.NewUserRepository)
	return server.UserServer{}
}
