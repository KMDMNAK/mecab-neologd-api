package parser

import (
	"fmt"
	"testing"
)

func TestExtractCountry(t *testing.T) {
	m, err := CreateMecabTagger(false)
	if err != nil {
		fmt.Println(err.Error())
	}
	var te = "ボーイング７３７貨物機が緊急着水 ハワイ・ホノルル沖"
	rns := m.GetNode(te)
	ae := AllExtractor{}
	for _, rn := range *rns {
		digger := CreateNodeDigger(rn, &ae)
		ns := digger.CreateNodesByRecursive()
		ce := CountryExtractor{}
		for _, n := range ns {
			if !ce.Filter(n) {
				continue
			}
			fmt.Println(n.Surface)
		}
	}
	t.Fail()
}

func TestExtractProper(t *testing.T) {
	m, err := CreateMecabTagger(true)
	if err != nil {
		fmt.Println(err.Error())
	}
	var te = "ボーイング７３７貨物機が緊急着水 ハワイ・ホノルル沖"
	rns := m.GetNode(te)
	ae := AllExtractor{}
	ce := ProperExtractor{}
	for _, rn := range *rns {
		digger := CreateNodeDigger(rn, &ae)
		ns := digger.CreateNodesByRecursive()
		for _, n := range ns {
			if !ce.Filter(n) {
				continue
			}
			fmt.Println(n.Surface)
		}
	}
	t.Fail()
}
