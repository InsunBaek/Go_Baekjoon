// Regular Expression
// https://www.acmicpc.net/problem/9342

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	exp, _ := regexp.Compile("^([ABCDEF]|)(A+F+C+)([ABCDEF]|)$")

	T := scanInt()
	for i := 0; i < T; i++ {
		if exp.MatchString(scanString()) {
			fmt.Fprintln(bw, "Infected!")
		} else {
			fmt.Fprintln(bw, "Good")
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
