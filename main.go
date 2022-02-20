package main

import (
	"epitime/database"
	routes "epitime/router"
	"github.com/gin-gonic/gin"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Cache-Control", "private, max-age=0")
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		//ctx.Header("Content-Encoding", "gzip")
		//ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT")
		//ctx.Writer.Header().Set("Content-Type", "application/json")

		ctx.Next()
	}
}

func main() {
	dba := database.NewEntDatabase()
	router := gin.Default()

	//router.Use(CORSMiddleware())
	routes.ApplyRoutes(router, dba)
	err := router.Run()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
