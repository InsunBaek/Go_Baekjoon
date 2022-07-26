// Implementation
// https://www.acmicpc.net/problem/9455

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func Fn() {
	Row, Col := scanInt(), scanInt()

	boxes := make([][]int, Row)
	for i := 0; i < Row; i++ {
		boxes[i] = make([]int, Col)
	}

	cnt := 0
	commuCol := make([]int, Col)

	for row := 0; row < Row; row++ {
		for col := 0; col < Col; col++ {
			boxes[row][col] = scanInt()
		}
	}

	for row := Row - 1; row >= 0; row-- {
		for col := 0; col < Col; col++ {
			if boxes[row][col] == 1 {
				cnt += ((row - Row + 1) + commuCol[col])
				commuCol[col]++
			}
		}
	}

	fmt.Fprintln(bw, -cnt)
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Tc := scanInt()

	for i := 0; i < Tc; i++ {
		Fn()
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
