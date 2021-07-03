package extract

import (
	"project.com/mecabapi/parser"
)

type Extractor interface {
	Filter(node *parser.Node) bool
}

type CountryExtractor struct{}

func (e *CountryExtractor) Filter(node *parser.Node) bool {
	return node.Attribute2 == "国"
}

type ProperExtractor struct{}

func (e *ProperExtractor) Filter(node *parser.Node) bool {
	return node.Kind == "固有名詞"
}
