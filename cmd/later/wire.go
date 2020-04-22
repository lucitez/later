//+build wireinject

package main

import (
	"database/sql"
	"later"
	"later/pkg/auth"
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
		service.NewContent,
		service.NewDomain,
		repository.NewContent,
		repository.NewDomain,
	)
	return server.Content{}
}

func InitializeDomain(db *sql.DB) server.Domain {
	wire.Build(
		server.NewDomain,
		service.NewDomain,
		repository.NewDomain,
	)
	return server.Domain{}
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
		service.NewContent,
		service.NewDomain,
		service.NewShare,
		service.NewUserContent,
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

func InitializeUserContent(db *sql.DB) server.UserContent {
	wire.Build(
		server.NewUserContent,
		service.NewUserContent,
		service.NewContent,
		service.NewDomain,
		service.NewUser,
		parse.NewContent,
		repository.NewUserContent,
		repository.NewContent,
		repository.NewDomain,
		repository.NewUser,
		transfer.NewUserContent,
	)
	return server.UserContent{}
}

func InitializeUser(db *sql.DB) server.User {
	wire.Build(
		server.NewUser,
		service.NewUser,
		service.NewFriendRequest,
		service.NewFriend,
		repository.NewUser,
		repository.NewFriendRequest,
		repository.NewFriend,
		transfer.NewUser,
	)
	return server.User{}
}

func InitializeAuth(db *sql.DB) server.Auth {
	wire.Build(
		server.NewAuth,
		service.NewUser,
		auth.NewService,
		repository.NewUser,
		repository.NewAuth,
	)
	return server.Auth{}
}

func InitializeServer(db *sql.DB) later.Server {
	wire.Build(
		later.NewServer,
		auth.NewService,
		repository.NewAuth,
	)
	return later.Server{}
}
