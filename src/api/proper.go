package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"project.com/mecabapi/extract"
	"project.com/mecabapi/parser"
)

type properHandler struct {
	pns []*parser.Node
}

func CreateProperHandler() *properHandler {
	ph := properHandler{}
	return &ph
}

type ProperRequestBody struct {
	Text string `form:"user" json:"text"`
}

func (e *properHandler) init(c *gin.Context) error {
	var text string
	var err error
	if c.Request.Method == "GET" {
		t := c.Param("text")
		text, err = url.QueryUnescape(t)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return err
		}
	} else if c.Request.Method == "POST" {
		var params ProperRequestBody
		err = c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return errors.New("不正な文字列です。")
		}
		text = params.Text
	}

	m, err := parser.CreateMecabTagger()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return err
	}
	rn, err := m.GetNode(text)
	if err != nil {
		c.JSON(http.StatusBadRequest, "不正なテキストです。")
		return err
	}
	ns := parser.CreateNodeByRecursive(rn)
	pe := extract.ProperExtractor{}
	pns := []*parser.Node{}
	for _, n := range ns {
		if !pe.Filter(n) {
			continue
		}
		pns = append(pns, n)
	}
	e.pns = pns
	return nil
}

func (e *properHandler) extractProper(c *gin.Context) {
	err := e.init(c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var ss = []string{}
	for _, p := range e.pns {
		ss = append(ss, p.Surface)
	}
	c.SecureJSON(http.StatusOK, map[string][]string{"proper_nouns": ss})
}

type ProperItem struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (e *properHandler) countProper(c *gin.Context) {
	err := e.init(c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var cmap = map[string]int{}
	for _, p := range e.pns {
		if _, ok := cmap[p.Surface]; ok {
			cmap[p.Surface]++
		} else {
			cmap[p.Surface] = 1
		}
	}
	items := []ProperItem{}
	for surface, count := range cmap {
		pi := ProperItem{
			Name:  surface,
			Count: count,
		}
		items = append(items, pi)
	}
	c.SecureJSON(http.StatusOK, map[string]interface{}{"proper_nouns_count": items})
}
