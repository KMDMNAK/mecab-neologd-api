package parser

import (
	"os"

	"github.com/shogo82148/go-mecab"
)

type MecabTagger struct {
	tagger mecab.MeCab
}

func (mt *MecabTagger) Destroy() {
	mt.tagger.Destroy()
}

func (mt *MecabTagger) GetNode(word string) (*mecab.Node, error) {
	node, err := mt.tagger.ParseToNode(word)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func CreateMecabTagger() (*MecabTagger, error) {
	np := os.Getenv("NEOLOGD_PATH")
	// http://www.mwsoft.jp/programming/munou/mecab_command.html
	args := map[string]string{
		"output-format-type": "chasen",
		// "output-format-type": "dump",
		"node-format": "%m,%s,%c\n",
		// "unk-format":  "%m,%m,%m,%f[6],%F-[0,1,2,3],,\n",
		// "eos-format":  "EOS,,,,,,\n",
	}
	if np != "" {
		args["dicdir"] = np
	}
	m, err := mecab.New(args)
	if err != nil {
		return nil, err
	}
	mt := MecabTagger{
		tagger: m,
	}
	return &mt, err
}
