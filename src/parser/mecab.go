package parser

import (
	"os"

	"github.com/shogo82148/go-mecab"
)

type MecabTagger struct {
	taggers []*mecab.MeCab
}

func (mt *MecabTagger) Destroy() {
	for _, tagger := range mt.taggers {
		tagger.Destroy()
	}
}

func (mt *MecabTagger) GetNode(word string) *[]*mecab.Node {
	var mns = []*mecab.Node{}
	for _, tagger := range mt.taggers {
		node, err := tagger.ParseToNode(word)
		if err != nil {
			continue
		}
		mns = append(mns, &node)
	}
	return &mns
}

type TaggerStock struct {
	taggers []*mecab.MeCab
}

func CreateMecabTagger(useNeologd bool) (*MecabTagger, error) {
	np := os.Getenv("NEOLOGD_PATH")
	// http://www.mwsoft.jp/programming/munou/mecab_command.html
	args := map[string]string{
		"output-format-type": "chasen",
		// "output-format-type": "wakati",
		// "output-format-type": "dump",
		"node-format": "%m,%s,%c\n",
		// "unk-format":  "%m,%m,%m,%f[6],%F-[0,1,2,3],,\n",
		// "eos-format":  "EOS,,,,,,\n",
	}
	mNormal, err := mecab.New(args)
	var mecabs []*mecab.MeCab
	if useNeologd {
		if np == "" {
			panic("NEOLOGD_PATHがありません。")
		}
		args["dicdir"] = np
		mNeo, err := mecab.New(args)
		if err != nil {
			panic(err.Error())
		}
		mecabs = append(mecabs, &mNeo)
	} else {
		mecabs = []*mecab.MeCab{&mNormal}
	}
	mt := MecabTagger{
		taggers: mecabs,
	}
	return &mt, err
}
