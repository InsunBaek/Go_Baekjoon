// Map
// https://www.acmicpc.net/problem/14425

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()

	br.Scan()
	token := strings.Fields(br.Text())

	N, _ := strconv.Atoi(token[0])
	M, _ := strconv.Atoi(token[1])
	m := make(map[string]bool)
	cnt := 0

	for i := 0; i < N; i++ {
		br.Scan()
		word := br.Text()
		m[word] = true
	}

	for i := 0; i < M; i++ {
		br.Scan()
		word := br.Text()
		if m[word] {
			cnt++
		}
	}

	bw.WriteString(strconv.Itoa(cnt))
}
