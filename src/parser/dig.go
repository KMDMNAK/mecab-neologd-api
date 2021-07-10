package parser

import (
	"fmt"

	"github.com/shogo82148/go-mecab"
)

type nodeDigger struct {
	rootNode *mecab.Node
	E        Extractor
}

func (nr *nodeDigger) CreateNodesByRecursive() []*Node {
	mnode := *nr.rootNode
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
			if nr.E.Filter(node) {
				mn = append(mn, node)
			}
		}
	}
	return mn
}

func CreateNodeDigger(
	rootNode *mecab.Node, E Extractor,
) *nodeDigger {
	d := nodeDigger{
		rootNode: rootNode,
		E:        E,
	}
	return &d
}
