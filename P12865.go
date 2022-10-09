// 01 Knapsack
// https://www.acmicpc.net/problem/12865

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Entity struct {
	value  int
	weight int
}

func Knapsack() {
	Entites := make([]Entity, 0)
	pcs, limit := scanInt()+1, scanInt()+1

	for i := 0; i < pcs; i++ {
		w, v := scanInt(), scanInt()
		Entites = append(Entites, Entity{value: v, weight: w})
	}

	DP := make([][]int, pcs)
	for i := 0; i < pcs; i++ {
		DP[i] = make([]int, limit)
	}
	for i := 0; i < pcs; i++ {
		DP[i][0] = 0
	}
	for i := 0; i < limit; i++ {
		DP[0][i] = 0
	}

	for n := 1; n < pcs; n++ {
		for w := 1; w < limit; w++ {
			if w-Entites[n-1].weight < 0 {
				DP[n][w] = DP[n-1][w]
			} else {
				DP[n][w] = Max(DP[n-1][w], (DP[n-1][w-Entites[n-1].weight] + Entites[n-1].value))
			}
		}
	}

	fmt.Fprintln(bw, DP[pcs-1][limit-1])
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Knapsack()
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
