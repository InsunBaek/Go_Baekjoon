// Implementation
// https://www.acmicpc.net/problem/11637

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	n := scanInt()
	total, major, ind, dup := 0, 0, 0, false

	for i := 0; i < n; i++ {
		val := scanInt()
		total += val
		if val > major {
			major = val
			ind = i + 1
			dup = false
		} else if val == major {
			dup = true
		}
	}

	if dup {
		fmt.Fprintln(bw, "no winner")
	} else {
		if major > (total / 2) {
			fmt.Fprintln(bw, "majority winner", ind)
		} else {
			fmt.Fprintln(bw, "minority winnder", ind)
		}
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Tc := scanInt()
	for Tc > 0 {
		Fn()
		Tc--
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
