package main

import (
	"fmt"
)

// 与えられた文字列の各文字を，以下の仕様で変換する関数cipherを実装せよ．
// 英小文字ならば(219 - 文字コード)の文字に置換
// その他の文字はそのまま出力
//
// この関数を用い，英語のメッセージを暗号化・復号化せよ．

func encrypt(s string) string {
	runes := []rune(s)
	c := make([]rune, len(runes), len(runes))
	for i := 0; i < len(runes); i++ {
		if 97 <= runes[i] && runes[i] <= 122 {
			c[i] = 219 - runes[i]
		} else {
			c[i] = runes[i]
		}
	}
	return string(c)
}

func decrypt(s string) string {
	return encrypt(s)
}

func main() {
	fmt.Printf("%+v\n", encrypt(encrypt("abCd")))
	fmt.Printf("%+v\n", decrypt(encrypt("abCd")))
}
