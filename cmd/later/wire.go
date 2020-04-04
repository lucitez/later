//+build wireinject

package main

import (
	"database/sql"
	"later/pkg/parse"
	"later/pkg/server"
	"later/pkg/service"
	"later/pkg/transfer"

	"later/pkg/repository"

	"github.com/google/wire"
)

func InitializeContent(db *sql.DB) server.Content {
	wire.Build(
		server.NewContent,
		parse.NewContent,
		service.NewContentManager,
		service.NewDomainManager,
		repository.NewContent,
		repository.NewDomain,
	)
	return server.Content{}
}

func InitializeDomain(db *sql.DB) server.DomainServer {
	wire.Build(
		server.NewDomainServer,
		service.NewDomainManager,
		repository.NewDomain,
	)
	return server.DomainServer{}
}

func InitializeFriend(db *sql.DB) server.Friend {
	wire.Build(
		server.NewFriend,
		service.NewFriend,
		service.NewUser,
		repository.NewFriend,
		repository.NewUser,
		transfer.NewFriend,
	)
	return server.Friend{}
}

func InitializeFriendRequest(db *sql.DB) server.FriendRequest {
	wire.Build(
		server.NewFriendRequest,
		service.NewFriendRequest,
		service.NewFriend,
		service.NewUser,
		repository.NewFriend,
		repository.NewFriendRequest,
		repository.NewUser,
		transfer.NewFriendRequest,
	)
	return server.FriendRequest{}
}

func InitializeShare(db *sql.DB) server.ShareServer {
	wire.Build(
		server.NewShareServer,
		service.NewContentManager,
		service.NewDomainManager,
		service.NewShare,
		service.NewUserContentManager,
		service.NewUser,
		repository.NewContent,
		repository.NewDomain,
		repository.NewShare,
		repository.NewUserContent,
		repository.NewUser,
		parse.NewContent,
	)
	return server.ShareServer{}
}

func InitializeUserContent(db *sql.DB) server.UserContentServer {
	wire.Build(
		server.NewUserContentServer,
		service.NewUserContentManager,
		repository.NewUserContent,
	)
	return server.UserContentServer{}
}

func InitializeUser(db *sql.DB) server.User {
	wire.Build(
		server.NewUser,
		service.NewUser,
		repository.NewUser,
	)
	return server.User{}
}
