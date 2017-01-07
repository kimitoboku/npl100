package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// スペースで区切られた単語列に対して，各単語の先頭と末尾の文字は残し，それ以外の文字の順序をランダムに並び替えるプログラムを作成せよ．ただし，長さが４以下の単語は並び替えないこととする．適当な英語の文（例えば"I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."）を与え，その実行結果を確認せよ．

func shuffle(s string) string {
	runes := []rune(s)
	for i := range s {
		j := rand.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func typoglycemia(s string) string {
	runes := []rune(s)
	if len(runes) > 5 {
		top := string(runes[0])
		end := string(runes[len(runes)-1])
		middle := shuffle(string(runes[1 : len(runes)-1]))
		return fmt.Sprintf("%s%s%s", top, middle, end)
	} else {
		return s
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	fmt.Printf("%+v\n", s)
	slist := strings.Split(s, " ")
	for i := range slist {
		fmt.Printf("%+v ", typoglycemia(slist[i]))
	}
}
