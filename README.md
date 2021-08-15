# pfv

PSDToolFavorites file parser

## Usage
```go
import (
	"github.com/raa0121/pvf"
)

func main() {
	s := `
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
	data := pvf.Decode(s)
	for _, v := range data.Items {
		fmt.Println(v.Name)
		for _, vv := range v.Elements {
			fmt.Printf("  %s\n", vv.Path)
		}
	}
}
```

## Requirements
go 1.16 or later


## Installation
```
go get github.com/raa0121/pfv
```

## License
Apache License 2.0

## Author
raa0121 <raa0121@gmail.com>
