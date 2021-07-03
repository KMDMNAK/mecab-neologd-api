package extract

import (
	"fmt"
	"testing"

	"project.com/mecabapi/parser"
)

func TestExtractCountry(t *testing.T) {
	m, err := parser.CreateMecabTagger()
	if err != nil {
		fmt.Println(err.Error())
	}
	var te = "ボーイング７３７貨物機が緊急着水 ハワイ・ホノルル沖"
	rn, _ := m.GetNode(te)
	ns := parser.CreateNodeByRecursive(rn)
	ce := CountryExtractor{}
	for _, n := range ns {
		if !ce.Filter(n) {
			continue
		}
		fmt.Println(n.Surface)
	}
	t.Fail()
}

func TestExtractProper(t *testing.T) {
	m, err := parser.CreateMecabTagger()
	if err != nil {
		fmt.Println(err.Error())
	}
	var te = "ボーイング７３７貨物機が緊急着水 ハワイ・ホノルル沖"
	rn, _ := m.GetNode(te)
	ns := parser.CreateNodeByRecursive(rn)
	ce := ProperExtractor{}
	for _, n := range ns {
		if !ce.Filter(n) {
			continue
		}
		fmt.Println(n.Surface)
	}
	t.Fail()
}
