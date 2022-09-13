// Math
// https://www.acmicpc.net/problem/2740

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, M int

func Fn() {

	N, M = scanInt(), scanInt()
	mA := make([][]int, N)
	for i := 0; i < N; i++ {
		mA[i] = make([]int, M)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			mA[i][j] = scanInt()
		}
	}

	M, K := scanInt(), scanInt()
	mB := make([][]int, M)
	for i := 0; i < M; i++ {
		mB[i] = make([]int, K)
	}
	for i := 0; i < M; i++ {
		for j := 0; j < K; j++ {
			mB[i][j] = scanInt()
		}
	}

	mC := make([][]int, N)
	for i := 0; i < N; i++ {
		mC[i] = make([]int, K)
	}

	for row := 0; row < N; row++ {
		for col := 0; col < K; col++ {
			for adr := 0; adr < M; adr++ {
				mC[row][col] += (mA[row][adr] * mB[adr][col])
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			fmt.Print(mC[i][j], " ")
		}
		fmt.Println()
	}

}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn()
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
