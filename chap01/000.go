package main

// 文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．

import (
	"fmt"
)

func main() {
	s := "stressed"
	fmt.Printf("%+v\n", s)

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	revs := string(runes)
	fmt.Printf("%+v\n", revs)
}
