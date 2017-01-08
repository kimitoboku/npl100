package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 行数をカウントせよ．確認にはwcコマンドを用いよ．

func main() {
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	defer fp.Close()
	linen := 0
	reader := bufio.NewReader(fp)
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			os.Exit(-1)
		} else {
			linen++
		}
	}
	fmt.Println(linen)
}
