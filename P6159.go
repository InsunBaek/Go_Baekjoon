// Sorting
// https://www.acmicpc.net/problem/6159

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var rd = bufio.NewReader(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, S := scanInt(), scanInt()
	intSlice := make([]int, N)
	cnt := 0

	for i := 0; i < N; i++ {
		intSlice[i] = scanInt()
	}
	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i] < intSlice[j]
	})
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if intSlice[i]+intSlice[j] <= S {
				cnt++
			}
		}
	}
	fmt.Fprintln(bw, cnt)
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
