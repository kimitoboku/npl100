package main

import (
	"fmt"
)

// 「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ．

func main() {
	s1 := "パトカー"
	s2 := "タクシー"
	rune1 := []rune(s1)
	rune2 := []rune(s2)
	comp := ""
	for i := 0; i < len(rune1); i++ {
		comp += string(rune1[i]) + string(rune2[i])
	}
	fmt.Printf("%+v\n", comp)

}
