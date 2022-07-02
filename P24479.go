// Graph, DFS
// https://www.acmicpc.net/problem/24479

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
var index []int
var graph []Edge
var visit []bool

var cnt = 0

type Edge struct {
	toList []int
}

func DFS(vertex int) {
	visit[vertex] = true
	cnt++
	index[vertex] = cnt

	for i := 0; i < len(graph[vertex].toList); i++ {
		nextVtx := graph[vertex].toList[i]
		if !visit[nextVtx] {
			DFS(nextVtx)
		}
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	vertex, edge, start := scanInt(), scanInt(), scanInt()
	index = make([]int, vertex+1)
	visit = make([]bool, vertex+1)
	graph = make([]Edge, vertex+1)
	for i := 0; i <= vertex; i++ {
		graph[i] = Edge{toList: make([]int, 0)}
	}

	for i := 0; i < edge; i++ {
		u, v := scanInt(), scanInt()
		graph[u].toList = append(graph[u].toList, v)
		graph[v].toList = append(graph[v].toList, u)
	}

	for i := 0; i < len(graph); i++ {
		sort.Slice(graph[i].toList, func(a, b int) bool {
			return graph[i].toList[a] < graph[i].toList[b]
		})
	}
	DFS(start)
	for i := 1; i < len(index); i++ {
		fmt.Fprintln(bw, index[i])
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
