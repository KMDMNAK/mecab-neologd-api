package api

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type ProperRequestBody struct {
	Text string `form:"user" json:"text"`
}

func extractRequestText(c *gin.Context) {
	var text string
	var err error
	if c.Request.Method == "GET" {
		t := c.Param("text")
		text, err = url.QueryUnescape(t)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
	} else if c.Request.Method == "POST" {
		var params ProperRequestBody
		err = c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, "不正な文字列です。")
			return
		}
		text = params.Text
	}
	c.Set("text", text)
	c.Next()
}
