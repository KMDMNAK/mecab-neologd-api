package api

import "project.com/mecabapi/parser"

type ProperItem struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func CountByName(pns *[]*parser.Node) *[]ProperItem {
	var cmap = map[string]int{}
	for _, p := range *pns {
		if _, ok := cmap[p.Surface]; ok {
			cmap[p.Surface]++
		} else {
			cmap[p.Surface] = 1
		}
	}
	items := []ProperItem{}
	for surface, count := range cmap {
		pi := ProperItem{
			Name:  surface,
			Count: count,
		}
		items = append(items, pi)
	}
	return &items
}

func GetName(pns *[]*parser.Node) *[]string {
	var ss = []string{}
	for _, n := range *pns {
		ss = append(ss, n.Surface)
	}
	return &ss
}
