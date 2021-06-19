package mecab

import (
	"fmt"
	"github.com/bluele/mecab-golang"
	"os"
	"strings"
)

type Environment struct {
	NEOLOGD_PATH string
}

// var env *Environment

// func Env() Environment {
// 	env.NeologdPath = os.Getenv("NeologdPath")
// 	return &env
// }

func GetNode(word string) string {
	// var word string
	var np string = os.Getenv("NEOLOGD_PATH")
	m, err := mecab.New("-Ochasen -d " + np)
	if err != nil {
		panic(err)
	}
	defer m.Destroy()
	// fmt.Print("input keyword :")
	// fmt.Scan(&word)
	return parseToNode(m, word)
}

func parseToNode(m *mecab.MeCab, word string) string {
	var key string
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(word)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] != "助詞" && features[0] != "助動詞" {
			key = key + " " + features[7]
		}
		if node.Next() != nil {
			break
		}
	}
	fmt.Println(key)
	return node
}
