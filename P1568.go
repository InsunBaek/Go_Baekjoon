// https://www.acmicpc.net/problem/1568
// Math

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
	second, cnt := 1, 0

	for N > 0 {

		if N < second {
			second = 1
			N -= second
			second++
			cnt++
		} else {
			N -= second
			second++
			cnt++
		}
	}
	fmt.Fprintln(bw, cnt)
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
