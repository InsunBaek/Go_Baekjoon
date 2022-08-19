// PrefixSum
// https://www.acmicpc.net/problem/17390

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	N, Q := scanInt(), scanInt()
	sequence := make([]int, 0)
	prefixSum := make([]int, N)

	for i := 0; i < N; i++ {
		sequence = append(sequence, scanInt())
	}
	sort.Slice(sequence, func(i, j int) bool {
		return sequence[i] < sequence[j]
	})

	prefixSum[0] = sequence[0]
	for i := 1; i < N; i++ {
		prefixSum[i] = prefixSum[i-1] + sequence[i]
	}

	for i := 0; i < Q; i++ {
		L, R := scanInt(), scanInt()
		fmt.Fprintln(bw, prefixSum[R-1]-prefixSum[L-1]+sequence[L-1])
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
