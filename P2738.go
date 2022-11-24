// Implementation
// https://www.acmicpc.net/problem/2738

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

	N, M := scanInt(), scanInt()

	m1, m2 := make([][]int, N), make([][]int, N)
	for i := 0; i < N; i++ {
		m1[i] = make([]int, M)
		m2[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			m1[i][j] = scanInt()
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			m2[i][j] = scanInt()
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			m1[i][j] += m2[i][j]
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Print(m1[i][j], " ")
		}
		fmt.Println()
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
