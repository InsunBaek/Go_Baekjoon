// https://www.acmicpc.net/problem/20949
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

type Monitor struct {
	PPI int
	num int
}

func Fn() {
	N := scanInt()
	slice := make([]Monitor, 0)

	for i := 0; i < N; i++ {
		W, H := scanInt(), scanInt()
		slice = append(slice, Monitor{PPI: ((W * W) + (H * H)), num: i + 1})
	}
	sort.Slice(slice, func(i, j int) bool {
		if slice[i].PPI == slice[j].PPI {
			return slice[i].num < slice[j].num
		} else {
			return slice[i].PPI > slice[j].PPI
		}
	})

	for i := 0; i < len(slice); i++ {
		fmt.Fprintln(bw, slice[i].num)
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
