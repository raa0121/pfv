package pfv

import (
	"testing"
	"reflect"
)

func TestDecode(t *testing.T) {
	data := `
[PSDToolFavorites-v1]
root-name/お気に入り
faview-mode/1

//Heart
マーク/ハート
アクセサリ/オタマン帽普
眉/*普通
目/*普通
口/*笑い
本体/*普段着
手/*手
ツインテール/*通常

//hoge/Heart2
マーク/ハート
アクセサリ/オタマンベレー帽
眉/*普通
目/*普通
口/*笑い
本体/*普段着
手/*手
ツインテール/*通常

`
	want := &Pfv{}
	want.Header = "[PSDToolFavorites-v1]"
	want.RootName = "お気に入り"
	want.FaviewMode = 1

	item := &Item{}
	item.Name = "Heart"
	elements := []*Element{}
	e := []string{
		"マーク/ハート",
		"アクセサリ/オタマン帽普",
		"眉/*普通",
		"目/*普通",
		"口/*笑い",
		"本体/*普段着",
		"手/*手",
		"ツインテール/*通常",
	}
	for _, el := range e {
		element := &Element{}
		element.Path = el
		elements = append(elements, element)
	}

	item.Elements = elements
	want.Items = append(want.Items, item)

	item = &Item{}
	item.Name = "hoge/Heart2"
	item.Elements = []*Element{}

	elements = []*Element{}
	e = []string{
		"マーク/ハート",
		"アクセサリ/オタマンベレー帽",
		"眉/*普通",
		"目/*普通",
		"口/*笑い",
		"本体/*普段着",
		"手/*手",
		"ツインテール/*通常",
	}
	for _, el := range e {
		element := &Element{}
		element.Path = el
		elements = append(elements, element)
	}

	item.Elements = elements
	want.Items = append(want.Items, item)

	t.Run("data", func(t *testing.T) {
		if got := Decode(data); !reflect.DeepEqual(got, want) {
			for i, _ := range got.Items {
				t.Errorf("Decode().Item[i] = %+v, want.Item %+v", got.Items[i], want.Items[i])
			}
		}
	})
}
