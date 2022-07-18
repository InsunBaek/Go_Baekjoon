// DFS, BFS
// https://www.acmicpc.net/problem/1325

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)
var cnt = 0
var visited []int

type Edge struct {
	toList []int
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()

	graph := make([]Edge, N+1)
	for i := 0; i <= N; i++ {
		graph[i].toList = make([]int, 0)
	}

	for i := 0; i < M; i++ {
		A, B := scanInt(), scanInt()
		graph[B].toList = append(graph[B].toList, A)
	}

	max := 0
	optSlice := make([]int, 0)
	for i := 1; i <= N; i++ {
		visited = make([]int, N+1)
		cnt = 0
		DFS(graph, i)

		if cnt > max {
			optSlice = make([]int, 0)
			optSlice = append(optSlice, i)
			max = cnt
		} else if cnt == max {
			optSlice = append(optSlice, i)
		}
	}

	sort.Slice(optSlice, func(i, j int) bool {
		return optSlice[i] < optSlice[j]
	})

	for i := 0; i < len(optSlice); i++ {
		fmt.Fprint(bw, optSlice[i], " ")
	}
}

func DFS(graph []Edge, now int) {
	visited[now] = 1

	for e := 0; e < len(graph[now].toList); e++ {
		nextVtx := graph[now].toList[e]
		if visited[nextVtx] == 0 {
			cnt++
			DFS(graph, nextVtx)
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
