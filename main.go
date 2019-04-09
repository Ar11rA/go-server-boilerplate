package main

import (
	"go/go-server-boilerplate/config"
	"go/go-server-boilerplate/routes"
)

func main() {
	router := routes.Router()
	s := config.Server{}
	s.Initialize(router)
	s.Run()
}
