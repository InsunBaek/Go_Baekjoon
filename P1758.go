// https://www.acmicpc.net/problem/1758
// Greedy

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

	N := scanInt()
	tip := make([]int, 0)

	for i := 0; i < N; i++ {
		tip = append(tip, scanInt())
	}
	sort.Slice(tip, func(i, j int) bool {
		return tip[i] > tip[j]
	})

	sum, order := 0, 1

	for i := 0; i < N; i++ {
		val := tip[i] - (order - 1)
		if val <= 0 {
			break
		} else {
			sum += val
		}
		order++
	}

	fmt.Fprintln(bw, sum)

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
