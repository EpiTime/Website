package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
)

func addIt(hideModules []string, str string) []string {
	hideModules = append(hideModules, str)
	return hideModules
}

func supIt(hideModules []string, i int) []string {
	hideModules = append(hideModules[:i], hideModules[(i+1):]...)
	return hideModules
}

func manageModules(sessionInter interface{}, str string) []string {
	hideModules := sessionInter.([]string)
	for i, v := range hideModules {
		if v == str {
			return supIt(hideModules, i)
		}
	}
	return addIt(hideModules, str)
}
func Modules(dba database) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		mod := c.Param("mod")
		sessionInter := session.Get("modules-hide")
		if sessionInter == nil {
			session.Set("modules-hide", []string{mod})
			err := session.Save()
			if err != nil {
				return
			}
			return
		}
		modArray := manageModules(sessionInter, mod)
		session.Set("modules-hide", modArray)
		email := session.Get("email")
		if email != nil {
			str3 := strings.Join(modArray, ",")
			err := dba.UpdateUserHideModules(c.Request.Context(), email.(string), str3)
			if err != nil {
				return
			}
		}
		err := session.Save()
		if err != nil {
			return
		}
	}
}
func GetModules(c *gin.Context) {
	session := sessions.Default(c)
	c.JSON(200, session.Get("modules-hide"))
}
