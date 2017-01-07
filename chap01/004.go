package main

import (
	"fmt"
	"strings"
)

// "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."という文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭に2文字を取り出し，取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成せよ．

func main() {
	s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	fmt.Printf("%+v\n", s)
	s = strings.Replace(s, ".", "", -1)
	slist := strings.Split(s, " ")
	indexList := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	elemDict := make(map[string]int)
	j := 0
	for i := 0; i < len(slist); i++ {
		if i+1 == indexList[j] {
			elemDict[slist[i][0:1]] = i + 1
			if j < len(indexList)-1 {
				j++
			}
		} else {
			elemDict[slist[i][0:2]] = i + 1
		}
	}
	fmt.Printf("%+v\n", elemDict)
}
