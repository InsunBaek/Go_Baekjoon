// DP
// https://www.acmicpc.net/problem/1003

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)
var DP []int

func Fn(n int) int {

	if n <= 0 {
		DP[0] = 0
		return 0
	} else if n == 1 {
		DP[1] = 1
		return 1
	}

	if DP[n] != 0 {
		return DP[n]
	} else {
		DP[n] = Fn(n-1) + Fn(n-2)
		return DP[n]
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Tc := scanInt()
	DP = make([]int, 41)

	for i := 0; i < Tc; i++ {
		N := scanInt()

		if N == 0 {
			fmt.Fprintln(bw, 1, 0)
		} else if N == 1 {
			fmt.Fprintln(bw, 0, 1)
		} else {
			Fn(N)
			fmt.Fprintln(bw, DP[N-1], DP[N])
			DP = make([]int, 41)
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
