package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 各行の1列目だけを抜き出したものをcol1.txtに，2列目だけを抜き出したものをcol2.txtとしてファイルに保存せよ．確認にはcutコマンドを用いよ．

func main() {
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)

	out1, err := os.Create("./../data/col1.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	defer out1.Close()
	out1Writer := bufio.NewWriter(out1)

	out2, err := os.Create("./../data/col2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	defer out2.Close()
	out2Writer := bufio.NewWriter(out2)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			os.Exit(-1)
		} else {
			s := strings.Replace(string(line), "\t", "\n\t", -1)
			ss := strings.Split(s, "\t")
			out1Writer.WriteString(ss[0])
			out2Writer.WriteString(ss[1])
		}
	}
	out1Writer.Flush()
	out2Writer.Flush()
}
