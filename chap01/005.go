package main

import (
	"fmt"
	"strings"
)

// 与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．この関数を用い，"I am an NLPer"という文から単語bi-gram，文字bi-gramを得よ．

func n_gram_ss(a []string, n int) [][]string {
	var list [][]string
	for i := 0; i < len(a)-n+1; i++ {
		list = append(list, a[i:i+n])
	}
	return list
}

func n_gram_s(s string, n int) []string {
	a := []rune(s)
	var list []string
	for i := 0; i < len(a)-n+1; i++ {
		list = append(list, string(a[i:i+n]))
	}
	return list
}

func main() {
	s := "I am an NLPer"
	ss := strings.Split(s, " ")
	fmt.Printf("%+v\n", n_gram_ss(ss, 2))
	fmt.Printf("%+v\n", n_gram_s(s, 2))

}
