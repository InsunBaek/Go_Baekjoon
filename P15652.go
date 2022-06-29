// Backtracking
// https://www.acmicpc.net/problem/15652

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
var arr []string

func nmFuntion(n int, m int, num int, cnt int) {
	if cnt == m {
		for i := 0; i < m; i++ {
			bw.WriteString(arr[i] + " ")
		}
		bw.WriteString("\n")
		return
	}
	for i := num; i <= n; i++ {
		arr[cnt] = strconv.Itoa(i)
		nmFuntion(n, m, i, cnt+1)
	}
}

func main() {
	defer bw.Flush()

	br.Scan()
	tokens := strings.Fields(br.Text())
	N, _ := strconv.Atoi(tokens[0])
	M, _ := strconv.Atoi(tokens[1])
	arr = make([]string, M)
	nmFuntion(N, M, 1, 0)
}
