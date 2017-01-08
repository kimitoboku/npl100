package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 行数をカウントせよ．確認にはwcコマンドを用いよ．

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
		fmt.Fprintf(os.Stderr, "Pleace input N > 0 : %v\n", n)
		os.Exit(-1)
	}

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
	i := 0
	fp.Seek(0, 0)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			os.Exit(-1)
		} else {
			if linen-int(n) <= i {
				fmt.Println(string(line))
			}
			i++
		}
	}

}
