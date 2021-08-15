package pfv

import (
	"strconv"
	"strings"
)

type Pfv struct {
	Header string
	RootName string
	FaviewMode int
	Items []*Item
}

type Item struct {
	Name string
	Elements []*Element
}

type Element struct {
	Path string
}

func Decode(s string) *Pfv {
	pfv := &Pfv{}
	data := strings.Split(s, "\n")
	for i, v := range data {
		if v == "[PSDToolFavorites-v1]" {
			pfv.Header = v
		}
		if strings.HasPrefix(v, "root-name/") {
			pfv.RootName = strings.Trim(v, "root-name/")
		}
		if strings.HasPrefix(v, "faview-mode/") {
			i, err := strconv.Atoi(strings.Trim(v, "faview-mode/"))
			if err != nil {
				panic(err)
			}
			pfv.FaviewMode = i
		}
		if strings.HasPrefix(v, `//`) {
			item := &Item{
				Name: strings.Trim(v, `//`),
			}
			for j, vv := range data[i+1:] {
				if vv == "" {
					i += j
					break
				}
				e := &Element{ Path: vv }
				item.Elements = append(item.Elements, e)
			}
			pfv.Items = append(pfv.Items, item)
		}
	}
	return pfv
}
