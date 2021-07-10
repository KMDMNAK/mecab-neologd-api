package parser

import (
	"errors"
	"strings"

	"github.com/shogo82148/go-mecab"
)

type Node struct {
	Surface       string `json:"word"`
	PartsOfSpeech string `json:"parts"`
	Kind          string `json:"kind"`
	Attribute1    string `json:"attribute1"`
	Attribute2    string `json:"attribute2"`
	Attribute3    string `json:"attribute3"`
	Attribute4    string `json:"attribute4"`
	Attribute5    string `json:"attribute5"`
	Attribute6    string `json:"attribute6"`
	Attribute7    string `json:"attribute7"`
}

func createNode(node *mecab.Node) (*Node, error) {
	feature := node.Feature()
	features := strings.Split(feature, ",")
	if len(features) < 9 {
		return nil, errors.New("不正なノードです。:" + node.Surface())
	}
	return &Node{
		Surface:       node.Surface(),
		PartsOfSpeech: features[0],
		Kind:          features[1],
		Attribute1:    features[2],
		Attribute2:    features[3],
		Attribute3:    features[4],
		Attribute4:    features[5],
		Attribute5:    features[6],
		Attribute6:    features[7],
		Attribute7:    features[8],
	}, nil
}
