package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// タブ1文字につきスペース1文字に置換せよ．確認にはsedコマンド，trコマンド，もしくはexpandコマンドを用いよ．

func main() {
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			os.Exit(-1)
		} else {
			s := strings.Replace(string(line), "\t", " ", -1)
			fmt.Println(s)
		}
	}
}
