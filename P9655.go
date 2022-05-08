// DP
// https://www.acmicpc.net/problem/9655

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var DP [1001]int

	br.Scan()
	N, _ := strconv.Atoi(br.Text())

	DP[0] = 0
	DP[1] = 1
	DP[2] = 2

	for i := 3; i <= N; i++ {
		DP[i] = int(math.Min(float64(DP[i-1]+1), float64(DP[i-3]+1)))
	}

	if DP[N]%2 == 1 {
		fmt.Println("SK")
	} else {
		fmt.Println("CY")
	}
}
