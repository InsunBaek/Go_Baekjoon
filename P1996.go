// Implementation
// https://www.acmicpc.net/problem/1996

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)
var dy = [...]int{1, 0, -1, 1, -1, 1, 0, -1}
var dx = [...]int{1, 1, 1, 0, 0, -1, -1, -1}

func canVisit(N int, row int, col int) bool {
	if row < 0 || col < 0 || row >= N || col >= N {
		return false
	}
	return true
}

func Fn() {
	N := scanInt()
	graph := make([][]string, N)
	makedG := make([][]int, N)

	for i := 0; i < N; i++ {
		graph[i] = make([]string, N)
		makedG[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		sc.Scan()
		line := sc.Text()
		for j := 0; j < N; j++ {
			graph[i][j] = string(line[j])
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if graph[i][j] != "." {
				for k := 0; k < 8; k++ {
					nRow := i + dy[k]
					nCol := j + dx[k]
					if canVisit(N, nRow, nCol) && graph[nRow][nCol] == "." {
						value, _ := strconv.Atoi(graph[i][j])
						makedG[nRow][nCol] += value
						if makedG[nRow][nCol] >= 10 {
							makedG[nRow][nCol] = math.MaxInt16
						}
					}
				}
				makedG[i][j] = math.MaxInt32
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if makedG[i][j] == math.MaxInt32 {
				fmt.Fprint(bw, "*")
			} else if makedG[i][j] == math.MaxInt16 {
				fmt.Fprint(bw, "M")
			} else {
				fmt.Fprint(bw, makedG[i][j])
			}
		}
		fmt.Fprintln(bw)
	}
}

func main() {
	defer bw.Flush()

	Fn()
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
