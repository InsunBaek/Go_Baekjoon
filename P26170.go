// Backtracking
// https://www.acmicpc.net/problem/26170

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

var info [5][5]int
var minValue int = math.MaxInt
var dy = [...]int{1, -1, 0, 0}
var dx = [...]int{0, 0, 1, -1}

func Fn() {
	var obstacle [5][5]bool

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			val := scanInt()
			if val == -1 {
				obstacle[i][j] = true
			}
			info[i][j] = val
		}
	}

	startRow, startCol := scanInt(), scanInt()

	DFS(obstacle, startRow, startCol, 0, 0)

	if minValue == math.MaxInt {
		fmt.Fprintln(bw, "-1")
	} else {
		fmt.Fprintln(bw, minValue)
	}
}

func DFS(obstacle [5][5]bool, row, col, cnt, try int) {
	obstacle[row][col] = true

	for i := 0; i < 4; i++ {
		newRow, newCol := row+dy[i], col+dx[i]
		if newRow < 0 || newRow > 4 || newCol < 0 || newCol > 4 {
			continue
		}
		if obstacle[newRow][newCol] || info[newRow][newCol] == -1 {
			continue
		}

		if info[newRow][newCol] == 1 {
			cnt++
			if cnt == 3 {
				if minValue > try+1 {
					minValue = try + 1
				}
				return
			}
			DFS(obstacle, newRow, newCol, cnt, try+1)
			cnt--
		} else {
			DFS(obstacle, newRow, newCol, cnt, try+1)
		}
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
