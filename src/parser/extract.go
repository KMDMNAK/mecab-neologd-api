package parser

type Extractor interface {
	Filter(*Node) bool
}

type CountryExtractor struct{}

func (e CountryExtractor) Filter(node *Node) bool {
	return node.Attribute2 == "国"
}

type GeoExtractor struct{}

func (e GeoExtractor) Filter(node *Node) bool {
	return node.Attribute1 == "地域" && node.Kind == "固有名詞"
}

type ProperExtractor struct{}

func (e ProperExtractor) Filter(node *Node) bool {
	return node.Kind == "固有名詞"
}

type AllExtractor struct{}

func (e AllExtractor) Filter(node *Node) bool {
	return true
}
