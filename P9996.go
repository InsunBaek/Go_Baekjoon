// Regular Expression
//https://www.acmicpc.net/problem/9996
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	N, pattern := scanInt(), scanString()
	var exp bytes.Buffer
	exp.WriteString("^")

	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 42 {
			exp.WriteString("[a-z]*")
		} else {
			exp.WriteString(string(pattern[i]))
		}
	}
	exp.WriteString("$")

	Exp, _ := regexp.Compile(exp.String())
	for i := 0; i < N; i++ {
		if Exp.MatchString(scanString()) {
			fmt.Fprintln(bw, "DA")
		} else {
			fmt.Fprintln(bw, "NE")
		}
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn()
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
