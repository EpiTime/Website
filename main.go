package main

import (
	"epitime/database"
	"epitime/router"
	"epitime/server"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	dba := database.NewEntDatabase()
	//user, err := create_user(context.Background(), dba)
	//fmt.Println(user)
	Engine := serverGest.NewServer()
	Engine.E.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.ApplyRoutes(Engine.E, dba)
	err := Engine.E.Run()
	if err != nil {
		return
	}
}
