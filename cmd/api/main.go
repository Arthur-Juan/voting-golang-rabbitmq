package main

import (
	"github.com/arthur-juan/voting-golang-rabbitmq/api/router"
	"github.com/arthur-juan/voting-golang-rabbitmq/config"
)

func main() {

	err := config.Init()
	if err != nil {
		panic(err)
	}

	router.Initialize()

}
