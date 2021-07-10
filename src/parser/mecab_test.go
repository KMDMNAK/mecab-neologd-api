package parser

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	m, err := CreateMecabTagger(false)
	defer m.Destroy()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}
	rns := m.GetNode(
		`
		UAE。
		ＵＡＥ。
		アラブ首長国連邦。
		アメリカから来ました。
		ボーイング７３７貨物機が緊急着水 ハワイ・ホノルル沖
		上田
		`)
	for _, rn := range *rns {
		node := *rn
		fmt.Println(node.Length())
		var count = 0
		for {
			node = node.Next()
			if node.Length() == 0 {
				break
			}
			fmt.Println(node.Surface())
			fmt.Println(node.Feature())
			count++
			if count > 100 {
				break
			}
		}

	}
	t.Fail()
}

func TestNode(t *testing.T) {
	m, err := CreateMecabTagger(true)
	defer m.Destroy()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}
	rns := m.GetNode("ふなっしー")
	for _, rn := range *rns {
		if err != nil {
			t.Log(err.Error())
			t.Fail()
			return
		}
		ae := AllExtractor{}
		digger := CreateNodeDigger(rn, &ae)
		nodes := digger.CreateNodesByRecursive()
		if len(nodes) != 1 {
			t.Log("辞書がNeologではありません。")
			t.Fail()
		}
		n := nodes[0]
		if n.Surface != "ふなっしー" || n.Attribute1 != "人名" || n.Kind != "固有名詞" || n.PartsOfSpeech != "名詞" {
			t.Log("辞書がNeologではありません。")
			t.Fail()
		}
	}

}
