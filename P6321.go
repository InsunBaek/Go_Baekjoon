// String
// https://www.acmicpc.net/problem/6321

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Change(word string) string {
	newWord := ""
	for i := 0; i < len(word); i++ {
		if word[i] == 90 {
			newWord += "A"
		} else {
			newWord += string(word[i] + 1)
		}
	}
	return newWord
}

func Fn() {
	N := scanInt()
	for i := 0; i < N; i++ {
		fmt.Fprint(bw, "String #")
		fmt.Fprintln(bw, i+1)
		fmt.Fprintln(bw, Change(scanString()))
		fmt.Fprintln(bw)
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn()
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}
