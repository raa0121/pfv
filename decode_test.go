package pfv

import (
	"testing"
	"reflect"
	"github.com/google/go-cmp/cmp"
)

func TestDecode1(t *testing.T) {
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
func TestDecode2(t *testing.T) {
	data := `
[PSDToolFavorites-v1]
root-name/riro
faview-mode/1

//zoi
!漫符/*なし
!顔パーツ/※ヘアピン
!顔パーツ/!眉/*＾＾
!顔パーツ/!口/*あ
!顔パーツ/!目/*普通目
!顔パーツ/!耳/*人間耳
!アホ毛/*普通
!表現/*なし
!おてて/右手/*ぞいの右手
!おてて/左手/*ぞいの左手
!前髪
!衣服/*通常服
!衣服/※タイツ
!衣服/!下着
!素体
ヘアースタイル/*ツインテール
※オタマン帽/!オタマン_目/*普通
※オタマン帽/!オタマン_目/!オタマン_目
※オタマン帽/!本体

//default
!漫符/*なし
!顔パーツ/※ヘアピン
!顔パーツ/!眉/*＾＾
!顔パーツ/!口/*あ
!顔パーツ/!目/*普通目
!顔パーツ/!耳/*人間耳
!アホ毛/*普通
!表現/*なし
!おてて/右手/*通常右手
!おてて/左手/*通常左手
!前髪
!衣服/*通常服
!衣服/※タイツ
!衣服/!下着
!素体
ヘアースタイル/*ツインテール
※オタマン帽/!オタマン_目/*普通
※オタマン帽/!オタマン_目/!オタマン_目
※オタマン帽/!本体

`
	want := &Pfv{}
	want.Header = "[PSDToolFavorites-v1]"
	want.RootName = "riro"
	want.FaviewMode = 1

	item := &Item{}
	item.Name = "zoi"
	elements := []*Element{}
	e := []string{
		"!漫符/*なし",
		"!顔パーツ/※ヘアピン",
		"!顔パーツ/!眉/*＾＾",
		"!顔パーツ/!口/*あ",
		"!顔パーツ/!目/*普通目",
		"!顔パーツ/!耳/*人間耳",
		"!アホ毛/*普通",
		"!表現/*なし",
		"!おてて/右手/*ぞいの右手",
		"!おてて/左手/*ぞいの左手",
		"!前髪",
		"!衣服/*通常服",
		"!衣服/※タイツ",
		"!衣服/!下着",
		"!素体",
		"ヘアースタイル/*ツインテール",
		"※オタマン帽/!オタマン_目/*普通",
		"※オタマン帽/!オタマン_目/!オタマン_目",
		"※オタマン帽/!本体",
	}
	for _, el := range e {
		element := &Element{}
		element.Path = el
		elements = append(elements, element)
	}

	item.Elements = elements
	want.Items = append(want.Items, item)

	item = &Item{}
	item.Name = "default"
	item.Elements = []*Element{}

	elements = []*Element{}
	e = []string{
		"!漫符/*なし",
		"!顔パーツ/※ヘアピン",
		"!顔パーツ/!眉/*＾＾",
		"!顔パーツ/!口/*あ",
		"!顔パーツ/!目/*普通目",
		"!顔パーツ/!耳/*人間耳",
		"!アホ毛/*普通",
		"!表現/*なし",
		"!おてて/右手/*通常右手",
		"!おてて/左手/*通常左手",
		"!前髪",
		"!衣服/*通常服",
		"!衣服/※タイツ",
		"!衣服/!下着",
		"!素体",
		"ヘアースタイル/*ツインテール",
		"※オタマン帽/!オタマン_目/*普通",
		"※オタマン帽/!オタマン_目/!オタマン_目",
		"※オタマン帽/!本体",
	}
	for _, el := range e {
		element := &Element{}
		element.Path = el
		elements = append(elements, element)
	}

	item.Elements = elements
	want.Items = append(want.Items, item)

	t.Run("data", func(t *testing.T) {
		got := Decode(data)
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("value is mismatch %s", diff)
		}
	})
}
