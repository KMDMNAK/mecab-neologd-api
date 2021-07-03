package api

import (
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

func (e *properHandler) init(c *gin.Context) {
	t := c.Param("text")
	text, err := url.QueryUnescape(t)
	m, err := parser.CreateMecabTagger()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	rn, err := m.GetNode(text)
	if err != nil {
		c.JSON(http.StatusBadRequest, "不正なテキストです。")
		return
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
}

func (e *properHandler) extractProper(c *gin.Context) {
	e.init(c)
	var ss = []string{}
	for _, p := range e.pns {
		ss = append(ss, p.Surface)
	}
	c.SecureJSON(http.StatusOK, map[string][]string{"proper_nouns": ss})
}

func (e *properHandler) countProper(c *gin.Context) {
	e.init(c)
	var cmap = map[string]int{}
	for _, p := range e.pns {
		if _, ok := cmap[p.Surface]; ok {
			cmap[p.Surface]++
		} else {
			cmap[p.Surface] = 1
		}
	}
	c.SecureJSON(http.StatusOK, map[string]interface{}{"proper_nouns_count": cmap})
}
