package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

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

type formatJson2 struct {
	Name   string        `json:"name"`
	Tag    string        `json:"tag"`
	Color  string        `json:"color"`
	Projet []formatJson3 `json:"projet"`
}

type FormatJson struct {
	Modules []formatJson2 `json:"modules"`
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
	year := session.Get("year").(string)
	if isValidYear(year) {
		file, err := ioutil.ReadFile(fmt.Sprintf("data/%s.json", year))
		if err != nil {
			return
		}
		var json1 = FormatJson{}
		err = json.Unmarshal(file, &json1)
		if err != nil {
			return
		}
		jsonTrie := SingleOut(json1, session)
		c.JSON(200, jsonTrie)
		return
	}
	c.String(400, "Year does not exist")
}

func Years(c *gin.Context) {
	year := c.Param("year")
	if isValidYear(year) {
		session := sessions.Default(c)
		session.Set("year", year)
		err := session.Save()
		if err != nil {
			return
		}
		c.Status(200)
		return
	}
	c.String(400, "Year does not exist")
}
