// https://www.acmicpc.net/problem/1448
// Sorting

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

func canTriangle(a, b, c int) bool {
	if a < b+c {
		return true
	}
	return false
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N := scanInt()
	straw := make([]int, 0)

	for i := 0; i < N; i++ {
		straw = append(straw, scanInt())
	}
	sort.Slice(straw, func(i, j int) bool {
		return straw[i] > straw[j]
	})

	length := 0

	for i := 0; i < len(straw)-2; i++ {
		a, b, c := straw[i], straw[i+1], straw[i+2]
		if canTriangle(a, b, c) {
			length = a + b + c
			break
		}
	}

	if length == 0 {
		fmt.Fprintln(bw, "-1")
	} else {
		fmt.Fprintln(bw, length)
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
