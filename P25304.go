// https://www.acmicpc.net/problem/25304
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

	Total, N := scanInt(), scanInt()
	total := 0

	for i := 0; i < N; i++ {
		price, n := scanInt(), scanInt()
		total += (price * n)
	}

	if Total == total {
		fmt.Fprintln(bw, "Yes")
	} else {
		fmt.Fprintln(bw, "No")
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
