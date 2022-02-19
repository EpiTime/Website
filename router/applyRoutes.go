package routes

import (
	"epitime/router/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(serv *gin.Engine) {
	store := cookie.NewStore([]byte("session"))
	serv.Use(sessions.Sessions("session", store))

	serv.GET("/health", routes.Health)
	serv.GET("/modules/:year", routes.Years)          // set la year dans les cookie
	serv.POST("/modules/toggle-diplay/:mod", routes.Modules) // met le module dans la liste des hide / retire le module de la liste des hides
	serv.GET("/modules/hidden", routes.GetModules)           // renvoie un array des modules hide
	serv.GET("/modules", routes.ShowTimeline)     // renvoie un json avec les modules de la year qui ne sont pas hide
}
