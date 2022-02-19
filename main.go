package main

import (
	"epitime/router"
	"epitime/server"
)

func main() {
	Engine := serverGest.NewServer()
	routes.ApplyRoutes(Engine.E)
	err := Engine.E.Run()
	if err != nil {
		return
	}
}
