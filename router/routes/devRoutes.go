package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.Status(200)
}

func GetMe(c *gin.Context) {
	s := sessions.Default(c)
	c.String(200, s.Get("email").(string))
}
