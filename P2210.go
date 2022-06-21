// Graph, DFS
// https://www.acmicpc.net/problem/2210

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
var dy = [...]int{1, -1, 0, 0}
var dx = [...]int{0, 0, 1, -1}
var graph = make([][]string, 5)

var m map[string]int
var CNT int = 0

func DFS(y int, x int, cnt int, str string) {
	if cnt == 6 {
		if m[str] == 0 {
			CNT++
			m[str] = 1
			return
		} else {
			return
		}
	}

	for i := 0; i < 4; i++ {
		ny := y + dy[i]
		nx := x + dx[i]

		if ny < 0 || nx < 0 || ny >= 5 || nx >= 5 {
			continue
		}
		DFS(ny, nx, cnt+1, str+graph[y][x])
	}
}

func main() {
	defer bw.Flush()
	m = make(map[string]int)

	for i := 0; i < len(graph); i++ {
		graph[i] = make([]string, 5)
	}

	for i := 0; i < 5; i++ {
		br.Scan()
		token := strings.Split(br.Text(), " ")

		for j := 0; j < 5; j++ {
			graph[i][j] = token[j]
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			DFS(i, j, 0, "")
		}
	}

	bw.WriteString(strconv.Itoa(CNT))
}
