package main

import (
	"epitime/database"
	routes "epitime/router"
	"github.com/gin-gonic/gin"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Content-Type", "application/json")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func main() {
	dba := database.NewEntDatabase()
	router := gin.Default()

	router.Use(CORSMiddleware())
	routes.ApplyRoutes(router, dba)
	err := router.Run()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
