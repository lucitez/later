package main

import (
	"fmt"
	"log"

	"later.co/pkg/parse"
	"later.co/pkg/repository"
	"later.co/pkg/repository/domainrepo"
)

func main() {

	db, err := repository.InitDB()

	if err != nil {
		log.Panic(err)
	}

	domainrepo.DB = db

	content, err := parse.ContentFromURL("https://soundcloud.com/itskforest/stwo-kmh-lovin-u")

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(content)
}
