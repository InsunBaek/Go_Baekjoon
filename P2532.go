// 별찍기
// https://www.acmicpc.net/problem/2523

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	N := scanInt()

	for i := 1; i <= N; i++ {
		for j := 0; j < i; j++ {
			fmt.Fprint(bw, "*")
		}
		fmt.Fprintln(bw)
	}
	for i := N - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Fprint(bw, "*")
		}
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
