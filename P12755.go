// Math
// https://www.acmicpc.net/problem/12755

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var rd = bufio.NewReader(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N := scanInt()
	start := 0
	flag := 0

	if N > (1 * 9) {
		N -= 9
		start = 10
		flag = 1
		if N > (2 * 90) {
			N -= (2 * 90)
			start = 100
			flag = 2
			if N > (3 * 900) {
				N -= (3 * 900)
				start = 1000
				flag = 3
				if N > (4 * 9000) {
					N -= (4 * 9000)
					start = 10000
					flag = 4
					if N > (5 * 90000) {
						N -= (5 * 90000)
						start = 100000
						flag = 5
						if N > (6 * 900000) {
							N -= (6 * 900000)
							start = 1000000
							flag = 6
							if N > (7 * 9000000) {
								N -= (7 * 9000000)
								start = 10000000
								flag = 7
							}
						}
					}
				}
			}
		}
	}

	unit := flag + 1
	n := N / unit
	m := N % unit

	if flag == 0 {
		fmt.Fprintln(bw, N)
	} else {
		if m == 0 {
			v := strconv.Itoa(start + n - 1)
			fmt.Fprintln(bw, string((v)[len(v)-1]))
		} else {
			v := strconv.Itoa(start + n)
			fmt.Fprintln(bw, string((v)[m-1]))
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
