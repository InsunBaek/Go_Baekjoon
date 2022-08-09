// https://www.acmicpc.net/problem/17266
// Implementation

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fn() {
	N := scanInt()
	M := scanInt()

	lightSlice := make([]int, 0)
	for i := 0; i < M; i++ {
		lightSlice = append(lightSlice, scanInt())
	}

	a1, a2 := lightSlice[0], N-lightSlice[M-1]
	min := Max(a1, a2)

	for i := 0; i < M-1; i++ {
		front, behind := lightSlice[i], lightSlice[i+1]
		gap := behind - front

		if gap%2 == 0 {
			min = Max(min, (gap / 2))
		} else {
			min = Max(min, ((gap + 1) / 2))
		}
	}

	fmt.Fprintln(bw, min)
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
