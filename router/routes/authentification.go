package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string
	Password string
}

var Users []User

func AddUser(email, password string) bool {
	for _, user := range Users {
		if user.Email == email {
			return false
		}
	}
	Users = append(Users, User{
		Email:    email,
		Password: password,
	})
	return true
}

func SignUp(c *gin.Context) {
	body := new(User)
	if err := c.ShouldBindJSON(body); err != nil {
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, body)
	AddUser(body.Email, body.Password)
	return
}

func SignIn(c *gin.Context) {
	body := new(User)
	if err := c.ShouldBindJSON(body); err != nil {
		c.AbortWithStatus(400)
		return
	}
	for _, v := range Users {
		if v.Email == body.Email && v.Password == body.Password {
			session := sessions.Default(c)
			session.Set("email", body.Email)
			err := session.Save()
			if err != nil {
				return
			}
			c.String(200, "CONNECTED")
			return
		}
	}
	c.String(400, "NO USERS")
}