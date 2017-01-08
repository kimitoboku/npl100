package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 自然数Nをコマンドライン引数などの手段で受け取り，入力のうち先頭のN行だけを表示せよ．確認にはheadコマンドを用いよ．

func main() {
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer fp.Close()
	n, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	} else if n < 1 {
		fmt.Fprintf(os.Stderr, "Place input N > 0:%v\n", n)
		os.Exit(-1)
	}
	linen := 0
	reader := bufio.NewReader(fp)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			os.Exit(-1)
		} else {
			fmt.Printf("%+v\n", string(line))
			linen++
			if int(n) == linen {
				break
			}
		}
	}
}
