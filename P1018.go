// Brute Force
// https://www.acmicpc.net/problem/1018

package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func minInt(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	defer bw.Flush()

	br.Scan()
	line := strings.Fields(br.Text())

	row, _ := strconv.Atoi(line[0])
	col, _ := strconv.Atoi(line[1])

	board := make([][]string, row)
	for i := 0; i < row; i++ {
		board[i] = make([]string, col)
	}

	for i := 0; i < row; i++ {
		br.Scan()
		block := br.Text()
		for j := 0; j < col; j++ {
			board[i][j] = string(block[j])
		}
	}

	rowLimit := row - 8
	colLimit := col - 8
	min := math.MaxInt

	for i := 0; i <= rowLimit; i++ {
		for j := 0; j <= colLimit; j++ {

			nRow := i
			nCol := j

			flag := true
			cnt_1 := 0
			cnt_2 := 0

			for r := nRow; r < nRow+8; r++ {
				for c := nCol; c < nCol+8; c++ {
					if flag {
						if board[r][c] == "B" {
							cnt_1++
						}
						if board[r][c] == "W" {
							cnt_2++
						}
						flag = false
					} else {
						if board[r][c] == "W" {
							cnt_1++
						}
						if board[r][c] == "B" {
							cnt_2++
						}
						flag = true
					}
				}
				flag = !flag
			}
			cnt := minInt(cnt_1, cnt_2)
			min = minInt(cnt, min)
		}
	}

	bw.WriteString(strconv.Itoa(min))
}
