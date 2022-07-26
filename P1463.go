// DP
// https://www.acmicpc.net/problem/1463

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func Fn() {
	N := scanInt()

	DP := make([]int, N+1)
	DP[1] = 0

	for i := 2; i <= N; i++ {
		DP[i] = DP[i-1] + 1
		if i%3 == 0 {
			DP[i] = intMin(DP[i], DP[i/3]+1)
		}
		if i%2 == 0 {
			DP[i] = intMin(DP[i], DP[i/2]+1)
		}
	}

	fmt.Fprintln(bw, DP[N])
}

func intMin(a int, b int) int {
	if a <= b {
		return a
	}
	return b
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
