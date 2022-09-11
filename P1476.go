// BruteForce
// https://www.acmicpc.net/problem/1476

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

	E, S, N := scanInt(), scanInt(), scanInt()
	e, s, n := 1, 1, 1
	date := 1

	for {
		if E == e && S == s && N == n {
			fmt.Println(date)
			break
		}

		if e < 15 {
			e++
		} else {
			e = 1
		}
		if s < 28 {
			s++
		} else {
			s = 1
		}
		if n < 19 {
			n++
		} else {
			n = 1
		}
		date++
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
