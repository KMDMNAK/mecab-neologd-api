package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"project.com/mecabapi/parser"
)

type handler struct {
	output func(pns []*parser.Node) interface{}
	m      *parser.MecabTagger
	E      parser.Extractor
}

func (h handler) handle(c *gin.Context) {
	text := c.MustGet("text").(string)
	nodes := h.m.GetNode(text)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	var nsArray = []*parser.Node{}
	for _, n := range *nodes {
		digger := parser.CreateNodeDigger(n, h.E)
		nsArray = append(nsArray, digger.CreateNodesByRecursive()...)
	}
	res := h.output(nsArray)
	c.SecureJSON(http.StatusOK, res)
}

func CreateExtractProperHandler(m *parser.MecabTagger) handler {
	h := handler{
		output: outputExtractProper,
		m:      m,
		E:      parser.ProperExtractor{},
	}
	return h
}

func outputExtractProper(pns []*parser.Node) interface{} {
	var ss = []string{}
	for _, n := range pns {
		ss = append(ss, n.Surface)
	}
	return map[string][]string{"proper_nouns": ss}
}

func CreateCountProperHandler(m *parser.MecabTagger) handler {
	h := handler{
		output: outputCountProper,
		m:      m,
		E:      parser.ProperExtractor{},
	}
	return h
}

func outputCountProper(pns []*parser.Node) interface{} {
	items := CountByName(&pns)
	return map[string]interface{}{"proper_nouns_count": *items}
}

func CreateExtractGeoHandler(m *parser.MecabTagger) handler {
	h := handler{
		output: outputExtractGeo,
		m:      m,
		E:      parser.GeoExtractor{},
	}
	return h
}

func outputExtractGeo(pns []*parser.Node) interface{} {
	ss := GetName(&pns)
	return map[string][]string{"geo": *ss}
}

func CreateExtractCountryHandler(m *parser.MecabTagger) handler {
	h := handler{
		output: outputExtractCountry,
		m:      m,
		E:      parser.CountryExtractor{},
	}
	return h
}

func outputExtractCountry(pns []*parser.Node) interface{} {
	ss := GetName(&pns)
	return map[string][]string{"countries": *ss}
}

func CreateExtractHandler(m *parser.MecabTagger) handler {
	h := handler{
		output: outputExtract,
		m:      m,
		E:      parser.AllExtractor{},
	}
	return h
}

func outputExtract(pns []*parser.Node) interface{} {
	var ss = []parser.Node{}
	for _, n := range pns {
		ss = append(ss, *n)
	}
	return map[string][]parser.Node{"words": ss}
}
