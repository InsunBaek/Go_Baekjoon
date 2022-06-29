// Greedy, Sort
// https://www.acmicpc.net/problem/11497

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func Abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func fn() int {
	N := scanInt()
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	n1 := arr[0]
	n2 := arr[1]
	max := Abs(n1, n2)

	for i := 2; i < N; i++ {
		if i%2 == 0 {
			max = Max(Abs(n1, arr[i]), max)
			n1 = arr[i]
		} else {
			max = Max(Abs(n2, arr[i]), max)
			n2 = arr[i]
		}
	}

	return max
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	tc := scanInt()
	for i := 0; i < tc; i++ {
		fmt.Fprintln(bw, fn())
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return int(v)
}
