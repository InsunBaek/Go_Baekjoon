// PrefixSum
// https://www.acmicpc.net/problem/11659

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)
var partial_sums [100001]int

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()

	for i := 1; i <= N; i++ {
		v := scanInt()
		partial_sums[i] = v + partial_sums[i-1]
	}

	for i := 0; i < M; i++ {
		i, j := scanInt(), scanInt()
		fmt.Fprintln(bw, partial_sums[j]-partial_sums[i-1])
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return int(v)
}
