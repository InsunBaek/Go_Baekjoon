// BruteForce
// https://www.acmicpc.net/problem/17086

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, M int

type Shark struct {
	Row int
	Col int
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, M = scanInt(), scanInt()
	Space := make([][]int, N)
	for i := 0; i < N; i++ {
		Space[i] = make([]int, M)
	}

	sharkSlice := make([]Shark, 0)
	for row := 0; row < N; row++ {
		for col := 0; col < M; col++ {
			v := scanInt()
			if v == 1 {
				sharkSlice = append(sharkSlice, Shark{Row: row, Col: col})
			} else {
				Space[row][col] = math.MaxInt
			}
		}
	}

	for row := 0; row < N; row++ {
		for col := 0; col < M; col++ {
			if Space[row][col] == math.MaxInt {
				for i := 0; i < len(sharkSlice); i++ {
					Space[row][col] = Min(Space[row][col], (Max(Abs(row-sharkSlice[i].Row), Abs(col-sharkSlice[i].Col))))
				}
			}
		}
	}

	max := 0
	for row := 0; row < N; row++ {
		for col := 0; col < M; col++ {
			if Space[row][col] > max {
				max = Space[row][col]
			}
		}
	}

	fmt.Fprintln(bw, max)
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
