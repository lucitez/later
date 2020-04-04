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
	domainManager := service.NewDomainManager(domain)
	content := parse.NewContent(domainManager)
	repositoryContent := repository.NewContent(db)
	contentManager := service.NewContentManager(repositoryContent)
	serverContent := server.NewContent(content, contentManager)
	return serverContent
}

func InitializeDomain(db *sql.DB) server.DomainServer {
	domain := repository.NewDomain(db)
	domainManager := service.NewDomainManager(domain)
	domainServer := server.NewDomainServer(domainManager)
	return domainServer
}

func InitializeFriend(db *sql.DB) server.Friend {
	user := repository.NewUser(db)
	serviceUser := service.NewUser(user)
	friend := repository.NewFriend(db)
	serviceFriend := service.NewFriend(serviceUser, friend)
	transferFriend := transfer.NewFriend(serviceUser)
	friendServer := server.NewFriend(serviceFriend, transferFriend)
	return friendServer
}

func InitializeFriendRequest(db *sql.DB) server.FriendRequest {
	friendRequest := repository.NewFriendRequest(db)
	user := repository.NewUser(db)
	serviceUser := service.NewUser(user)
	friend := repository.NewFriend(db)
	serviceFriend := service.NewFriend(serviceUser, friend)
	serviceFriendRequest := service.NewFriendRequest(friendRequest, serviceFriend, serviceUser)
	transferFriendRequest := transfer.NewFriendRequest(serviceUser)
	friendRequestServer := server.NewFriendRequest(serviceFriendRequest, transferFriendRequest)
	return friendRequestServer
}

func InitializeShare(db *sql.DB) server.ShareServer {
	share := repository.NewShare(db)
	userContent := repository.NewUserContent(db)
	userContentManager := service.NewUserContentManager(userContent)
	shareManager := service.NewShare(share, userContentManager)
	content := repository.NewContent(db)
	contentManager := service.NewContentManager(content)
	user := repository.NewUser(db)
	serviceUser := service.NewUser(user)
	domain := repository.NewDomain(db)
	domainManager := service.NewDomainManager(domain)
	parseContent := parse.NewContent(domainManager)
	shareServer := server.NewShareServer(shareManager, contentManager, serviceUser, parseContent)
	return shareServer
}

func InitializeUserContent(db *sql.DB) server.UserContentServer {
	userContent := repository.NewUserContent(db)
	userContentManager := service.NewUserContentManager(userContent)
	userContentServer := server.NewUserContentServer(userContentManager)
	return userContentServer
}

func InitializeUser(db *sql.DB) server.User {
	user := repository.NewUser(db)
	serviceUser := service.NewUser(user)
	userServer := server.NewUser(serviceUser)
	return userServer
}
