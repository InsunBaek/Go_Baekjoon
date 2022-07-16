// Sorting
// https://www.acmicpc.net/problem/11931

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

	N := scanInt()
	intSlice := make([]int, N)

	for i := 0; i < N; i++ {
		intSlice[i] = scanInt()
	}
	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i] > intSlice[j]
	})
	for i := 0; i < N; i++ {
		fmt.Fprintln(bw, intSlice[i])
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
