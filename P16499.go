// https://www.acmicpc.net/problem/16499
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

func Fn() {
	cToc := make(map[string]bool)

	N := scanInt()
	cnt := 0

	for i := 0; i < N; i++ {
		sc.Scan()
		word := sc.Text()

		slice := []rune(word)

		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})

		key := string(slice)
		if !cToc[key] {
			cToc[key] = true
			cnt++
		}
	}

	fmt.Fprintln(bw, cnt)
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
