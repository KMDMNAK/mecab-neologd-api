package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shogo82148/go-mecab"
)

type Node struct {
	Surface       string
	PartsOfSpeech string
	Kind          string
	Attribute1    string
	Attribute2    string
}

func createNode(node *mecab.Node) (*Node, error) {
	feature := node.Feature()
	features := strings.Split(feature, ",")
	if len(features) < 3 {
		return nil, errors.New("不正なノードです。")
	}
	return &Node{
		Surface:       node.Surface(),
		PartsOfSpeech: features[0],
		Kind:          features[1],
		Attribute1:    features[2],
		Attribute2:    features[3],
	}, nil
}

func CreateNodeByRecursive(rn *mecab.Node) []*Node {
	mnode := *rn
	mn := []*Node{}
	for {
		mnode = mnode.Next()
		if mnode.Length() == 0 {
			break
		}
		node, err := createNode(&mnode)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			mn = append(mn, node)
		}
	}
	return mn
}
