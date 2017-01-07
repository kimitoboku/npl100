package main

import (
	"fmt"
	"github.com/kimitoboku/npl100/chap01/set"
)

// "paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，それぞれ, XとYとして求め，XとYの和集合，積集合，差集合を求めよ．さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ．

func n_gram_s(s string, n int) []string {
	a := []rune(s)
	var list []string
	for i := 0; i < len(a)-n+1; i++ {
		list = append(list, string(a[i:i+n]))
	}
	return list
}

func main() {
	s1 := "paraparaparadise"
	s2 := "paragraph"
	set1 := set.NetSet()
	set2 := set.NetSet()
	set1.Add(n_gram_s(s1, 2)...)
	set2.Add(n_gram_s(s2, 2)...)
	fmt.Printf("%+v\n", set1)
	fmt.Printf("%+v\n", set2)
	fmt.Printf("%+v\n", set1.Union(set2))
	fmt.Printf("%+v\n", set1.Intersection(set2))
}
