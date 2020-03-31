package main

import (
	"fmt"
	"log"

	"later/pkg/parse"
	"later/pkg/repository"
	"later/pkg/repository/domainrepo"
)

func main() {

	db, err := repository.InitDB()

	if err != nil {
		log.Panic(err)
	}

	domainrepo.DB = db

	_, err = parse.ContentFromURL("https://soundcloud.com/itskforest/stwo-kmh-lovin-u")

	if err != nil {
		fmt.Print(err)
	}
}
