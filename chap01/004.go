package main

import (
	"fmt"
	"strings"
)

// "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．

func main() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	fmt.Printf("%+v\n", s)

	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	slist := strings.Split(s, " ")
	var piList []int
	for _, st := range slist {
		piList = append(piList, len(st))
	}
	fmt.Printf("%+v\n", piList)

}
