// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"database/sql"
	"later/pkg/parse"
	"later/pkg/repository"
	"later/pkg/server"
	"later/pkg/service"
	"later/pkg/transfer"
)

// Injectors from wire.go:

func InitializeContent(db *sql.DB) server.Content {
	domain := repository.NewDomain(db)
	serviceDomain := service.NewDomain(domain)
	content := repository.NewContent(db)
	parseContent := parse.NewContent()
	serviceContent := service.NewContent(serviceDomain, content, parseContent)
	serverContent := server.NewContent(serviceContent)
	return serverContent
}

func InitializeDomain(db *sql.DB) server.Domain {
	domain := repository.NewDomain(db)
	serviceDomain := service.NewDomain(domain)
	serverDomain := server.NewDomain(serviceDomain)
	return serverDomain
}

func InitializeFriend(db *sql.DB) server.Friend {
	user := repository.NewUser(db)
	serviceUser := service.NewUser(user)
	friend := repository.NewFriend(db)
	serviceFriend := service.NewFriend(serviceUser, friend)
	transferFriend := transfer.NewFriend(serviceUser)
	serverFriend := server.NewFriend(serviceFriend, transferFriend)
	return serverFriend
}

func InitializeFriendRequest(db *sql.DB) server.FriendRequest {
	friendRequest := repository.NewFriendRequest(db)
	user := repository.NewUser(db)
	serviceUser := service.NewUser(user)
	friend := repository.NewFriend(db)
	serviceFriend := service.NewFriend(serviceUser, friend)
	serviceFriendRequest := service.NewFriendRequest(friendRequest, serviceFriend, serviceUser)
	transferFriendRequest := transfer.NewFriendRequest(serviceUser)
	serverFriendRequest := server.NewFriendRequest(serviceFriendRequest, transferFriendRequest)
	return serverFriendRequest
}

func InitializeShare(db *sql.DB) server.ShareServer {
	share := repository.NewShare(db)
	userContent := repository.NewUserContent(db)
	serviceUserContent := service.NewUserContent(userContent)
	serviceShare := service.NewShare(share, serviceUserContent)
	domain := repository.NewDomain(db)
	serviceDomain := service.NewDomain(domain)
	content := repository.NewContent(db)
	parseContent := parse.NewContent()
	serviceContent := service.NewContent(serviceDomain, content, parseContent)
	user := repository.NewUser(db)
	serviceUser := service.NewUser(user)
	shareServer := server.NewShareServer(serviceShare, serviceContent, serviceUser, parseContent)
	return shareServer
}

func InitializeUserContent(db *sql.DB) server.UserContentServer {
	userContent := repository.NewUserContent(db)
	serviceUserContent := service.NewUserContent(userContent)
	userContentServer := server.NewUserContentServer(serviceUserContent)
	return userContentServer
}

func InitializeUser(db *sql.DB) server.User {
	user := repository.NewUser(db)
	serviceUser := service.NewUser(user)
	serverUser := server.NewUser(serviceUser)
	return serverUser
}
