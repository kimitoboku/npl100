package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 12で作ったcol1.txtとcol2.txtを結合し，元のファイルの1列目と2列目をタブ区切りで並べたテキストファイルを作成せよ．確認にはpasteコマンドを用いよ．

func main() {
	fp1, err := os.Open("../data/col1.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	defer fp1.Close()
	reader1 := bufio.NewReader(fp1)
	fp2, err := os.Open("../data/col2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	defer fp2.Close()
	reader2 := bufio.NewReader(fp2)

	marge, err := os.Create("./../data/marge.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	defer marge.Close()
	outWriter := bufio.NewWriter(marge)

	for {
		col1, _, err := reader1.ReadLine()
		col2, _, err := reader2.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			os.Exit(-1)
		} else {
			outWriter.WriteString(fmt.Sprintf("%v\t%v\n", string(col1), string(col2)))
		}
	}
	outWriter.Flush()
}
