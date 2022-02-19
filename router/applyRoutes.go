package routes

import (
	"context"
	"epitime/database"
	"epitime/router/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(serv *gin.Engine, dba database.Database) {
	store := cookie.NewStore([]byte("session"))
	serv.Use(sessions.Sessions("session", store))

	serv.GET("/health", routes.Health)
	serv.GET("/getme", routes.GetMe)
	serv.POST("/signUp", routes.SignItUp(context.Background(), dba.Client))
	serv.POST("/signIn", routes.SignItIn(context.Background(), dba.Client))
	serv.POST("/year/:year", routes.Years(dba))                    // set la year dans les cookie
	serv.POST("/modules/toggle-display/:mod", routes.Modules(dba)) // met le module dans la liste des hide / retire le module de la liste des hides
	serv.GET("/modules/hidden", routes.GetModules)                 // renvoie un array des modules hide
	serv.GET("/modules", routes.ShowTimeline)                      // renvoie un json avec les modules de la year qui ne sont pas hide
	serv.POST("/modules/add", routes.AddModules(dba))
}
