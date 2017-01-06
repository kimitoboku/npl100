package main

import (
	"fmt"
)

// 「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ．

func main() {
	s := "パタトクカシーー"
	fmt.Printf("%+v\n", s)
	runes := []rune(s)
	ext := ""
	for i := 0; i < len(runes); i++ {
		if i%2 == 0 {
			ext += string(runes[i])
		}
	}
	fmt.Printf("%+v\n", ext)
}
