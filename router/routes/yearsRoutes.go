package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type database interface {
	UpdateUserYear(ctx context.Context, email string, year int) error
	UpdateUserHideModules(ctx context.Context, email string, hideModules string) error
}

func isValidYear(year string) bool {
	var currentYears [5]string
	jsonString, _ := ioutil.ReadFile("data/currentYears.json")
	err := json.Unmarshal(jsonString, &currentYears)
	if err != nil {
		return false
	}
	for _, v := range currentYears {
		if year == v {
			return true
		}
	}
	return false
}

type formatJson3 struct {
	Name  string `json:"name"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type FormatJson2 struct {
	Name    string        `json:"name"`
	Tag     string        `json:"tag"`
	Color   string        `json:"color"`
	Project []formatJson3 `json:"project"`
}

type FormatJson struct {
	Modules []FormatJson2 `json:"modules"`
}

func isIn(listHide []string, mod string) bool {
	for _, v := range listHide {
		if v == mod {
			return true
		}
	}
	return false
}

func SingleOut(fullJson FormatJson, sess sessions.Session) FormatJson {
	sessionInter := sess.Get("modules-hide")
	if sessionInter == nil {
		return fullJson
	}
	listHide := sess.Get("modules-hide").([]string)
	var viewedJson FormatJson
	for _, v := range fullJson.Modules {
		if !isIn(listHide, v.Name) {
			viewedJson.Modules = append(viewedJson.Modules, v)
		}
	}
	return viewedJson
}

func ShowTimeline(c *gin.Context) {
	session := sessions.Default(c)
	if !(session.Get("year") != nil) {
		c.String(400, "Year not set up")
	}
	year := session.Get("year").(string)
	if !isValidYear(year) {
		c.String(400, "Year does not exist")
	}
	file, err := ioutil.ReadFile(fmt.Sprintf("data/%s.json", year))
	if err != nil {
		return
	}
	var json1 = FormatJson{}
	err = json.Unmarshal(file, &json1)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	jsonTrie := SingleOut(json1, session)
	rawJson, _ := json.Marshal(jsonTrie)
	jsonData := []byte(rawJson)

	c.Data(http.StatusOK, "application/json", jsonData)
	//c.JSON(200, jsonTrie)
	return

}

func Years(dba database) gin.HandlerFunc {
	return func(c *gin.Context) {
		year := c.Param("year")
		if !isValidYear(year) {
			c.String(400, "Year does not exist")
			return
		}
		yearInt, _ := strconv.Atoi(year)
		session := sessions.Default(c)
		session.Set("year", year)
		err := session.Save()
		email := session.Get("email")
		if email != nil {
			err := dba.UpdateUserYear(c.Request.Context(), email.(string), yearInt)
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
		c.Status(200)
		return
	}
}
