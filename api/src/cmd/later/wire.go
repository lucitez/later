//+build wireinject

package main

import (
	"database/sql"
	"later"
	"later/pkg/auth"
	"later/pkg/server"
	"later/pkg/service"
	"later/pkg/transfer"

	"later/pkg/repository"

	"github.com/google/wire"
)

func InitializeContent(db *sql.DB) server.Content {
	wire.Build(
		server.NewContent,
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
		service.NewChat,
		service.NewMessage,
		repository.NewContent,
		repository.NewDomain,
		repository.NewShare,
		repository.NewUserContent,
		repository.NewUser,
		repository.NewChat,
		repository.NewMessage,
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
		service.NewContent,
		service.NewDomain,
		repository.NewUser,
		repository.NewFriendRequest,
		repository.NewFriend,
		repository.NewContent,
		repository.NewDomain,
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
		repository.NewUser,
		repository.NewAuth,
	)
	return later.Server{}
}

func InitializeChat(db *sql.DB) server.Chat {
	wire.Build(
		service.NewUser,
		repository.NewUser,
		service.NewChat,
		service.NewMessage,
		repository.NewChat,
		repository.NewMessage,
		server.NewChat,
		transfer.NewChat,
	)

	return server.Chat{}
}

func InitializeMessage(db *sql.DB) server.Message {
	wire.Build(
		server.NewMessage,
		service.NewChat,
		service.NewMessage,
		service.NewContent,
		service.NewDomain,
		repository.NewChat,
		repository.NewMessage,
		repository.NewContent,
		repository.NewDomain,
		transfer.NewMessage,
	)

	return server.Message{}
}
