//+build wireinject

package main

import (
	"database/sql"
	"later/pkg/parse"
	"later/pkg/server"
	"later/pkg/service"

	"later/pkg/repository"

	"github.com/google/wire"
)

func InitializeContent(db *sql.DB) server.Content {
	wire.Build(
		server.NewContent,
		parse.NewContent,
		service.NewContentManager,
		service.NewDomainManager,
		repository.NewContentRepository,
		repository.NewDomainRepository)
	return server.Content{}
}

func InitializeDomain(db *sql.DB) server.DomainServer {
	wire.Build(
		server.NewDomainServer,
		service.NewDomainManager,
		repository.NewDomainRepository)
	return server.DomainServer{}
}

func InitializeFriend(db *sql.DB) server.FriendServer {
	wire.Build(
		server.NewFriendServer,
		service.NewFriendManager,
		service.NewUserManager,
		repository.NewFriend,
		repository.NewUser)
	return server.FriendServer{}
}

func InitializeFriendRequest(db *sql.DB) server.FriendRequestServer {
	wire.Build(
		server.NewFriendRequestServer,
		service.NewFriendRequestManager,
		repository.NewFriendRequestRepository)
	return server.FriendRequestServer{}
}

func InitializeShare(db *sql.DB) server.ShareServer {
	wire.Build(
		server.NewShareServer,
		service.NewContentManager,
		service.NewDomainManager,
		service.NewShareManager,
		service.NewUserContentManager,
		service.NewUserManager,
		repository.NewContentRepository,
		repository.NewDomainRepository,
		repository.NewShare,
		repository.NewUserContent,
		repository.NewUser,
		parse.NewContent)
	return server.ShareServer{}
}

func InitializeUserContent(db *sql.DB) server.UserContentServer {
	wire.Build(
		server.NewUserContentServer,
		service.NewUserContentManager,
		repository.NewUserContent)
	return server.UserContentServer{}
}

func InitializeUser(db *sql.DB) server.UserServer {
	wire.Build(
		server.NewUserServer,
		service.NewUserManager,
		repository.NewUser)
	return server.UserServer{}
}
