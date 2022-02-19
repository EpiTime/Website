package main

import (
	"epitime/database"
	"epitime/router"
	"epitime/server"
)

func main() {
	dba := database.Database{}
	dba.Client = database.NewEntDatabase(dba)
	//user, err := create_user(context.Background(), dba)
	//fmt.Println(user)
	Engine := serverGest.NewServer()
	routes.ApplyRoutes(Engine.E, dba)
	err := Engine.E.Run()
	if err != nil {
		return
	}
}
