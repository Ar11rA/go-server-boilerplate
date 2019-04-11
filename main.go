package main

import (
	"go/go-server-boilerplate/config"
)

func main() {
	s := config.Server{}
	s.InitializeDB()
	s.InitializeEntities()
	s.InitializeRoutes()
	s.Run()
}
