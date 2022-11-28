// Implementation
// https://www.acmicpc.net/problem/2566

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

	cRow, cCol, max := 0, 0, -1

	for row := 1; row <= 9; row++ {
		for col := 1; col <= 9; col++ {
			v := scanInt()
			if v > max {
				max = v
				cRow, cCol = row, col
			}
		}
	}
	fmt.Fprintln(bw, max)
	fmt.Fprintln(bw, cRow, cCol)
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
