package routes

import (
	"context"
	"epitime/ent"
	"epitime/ent/user"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserStruct struct {
	Email    string
	Password string
}

func setupUser(ctx *gin.Context, client *ent.Client) {

}

func createUser(ctx *gin.Context, client *ent.Client, email, password string) (*ent.User, error) {
	u, err := client.User.Create().
		SetEmail(email).
		SetPassword(password).
		Save(ctx)
	sess := sessions.Default(ctx)
	e := sess.Get("year")
	if e != nil {
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
	}
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return u, err
}

func SignItUp(ctx context.Context, client *ent.Client) gin.HandlerFunc {
	f := func(c *gin.Context) {
		u := new(UserStruct)
		if err := c.ShouldBindJSON(u); err != nil {
			c.AbortWithStatus(400)
			return
		}
		_, err := client.User.Query().Where(user.Email(u.Email)).All(c)
		if err == nil {
			_, err := createUser(c, client, u.Email, u.Password)
			c.String(200, "Successfully connected")
			if err != nil {
				return
			}
		}
		return
	}
	return f
}

func SignItIn(c context.Context, client *ent.Client) (f gin.HandlerFunc) {
	f = func(c *gin.Context) {
		u := new(UserStruct)
		if err := c.ShouldBindJSON(u); err != nil {
			c.AbortWithStatus(400)
			return
		}
		e, _ := client.User.Query().Where(user.Email(u.Email)).All(c)
		if e[0].Password == u.Password {
			session := sessions.Default(c)
			session.Set("email", u.Email)
			if len(e[0].HideModules) > 0 {
				session.Set("modules-hide", strings.Split(e[0].HideModules, ","))
			}
			if e[0].Year > 0 {
				ye := strconv.Itoa(e[0].Year)
				session.Set("year", ye)
			}
			err := session.Save()
			if err != nil {
				return
			}
			c.String(200, "Nice you connected")
			return
		}
		c.String(http.StatusUnauthorized, "Bad mdp")
	}
	return
}

//	u := dba.Client.Artist.Query().Where(artist.ID(uuid2.UUID(id))).AllX(ctx)
