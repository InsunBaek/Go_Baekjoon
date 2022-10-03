// MST
// https://www.acmicpc.net/problem/1197

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

var graph []map[int]int

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	pq := &EdgeHeap{}
	heap.Init(pq)

	V, E := scanInt(), scanInt()
	graph = make([]map[int]int, V+1)
	for i := 0; i <= V; i++ {
		graph[i] = make(map[int]int)
	}
	UFInit(V + 1)

	for i := 0; i < E; i++ {
		from, to, cost := scanInt(), scanInt(), scanInt()

		graph[from][to] = cost

		heap.Push(pq, &Edge{V1: from, V2: to, Cost: cost})
	}

	totalCost := 0

	for pq.Len() > 0 {
		myEdge := heap.Pop(pq).(Edge)
		v1, v2 := myEdge.V1, myEdge.V2

		p1, p2 := find(v1), find(v2)
		if p1 == p2 {
			continue
		}
		totalCost += myEdge.Cost
		union(v1, v2)
	}

	fmt.Println(totalCost)
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

/* Priority Queue */

type Edge struct {
	V1, V2 int
	Cost   int
}

type EdgeHeap []*Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(u interface{}) { *h = append(*h, u.(*Edge)) }
func (h *EdgeHeap) Pop() interface{} {
	old := *h
	long := len(old)
	x := old[long-1]
	*h = old[0 : long-1]
	return *x
}

/* UnionFind */

var parent []int

func UFInit(N int) {
	parent = make([]int, N)
	for i := 0; i < N; i++ {
		parent[i] = -1
	}
}

// weighting Rule
func union(x_, y_ int) {
	x := find(x_)
	y := find(y_)

	// root가 같다면 사이클 형성
	if x == y {
		return
	}

	// 적은 node를 leaf로 가진 node를 아래에 붙임
	if parent[x] < parent[y] {
		parent[x] = parent[x] + parent[y]
		parent[y] = x
	} else {
		parent[y] = parent[y] + parent[x]
		parent[x] = y
	}
}

// Collapsing Rule
func find(node int) int {
	var root, trail, front int
	// node(left)에서 출발해서 parent값이 음수인(뿌리)노드를 탐색
	for root = node; parent[root] >= 0; {
		root = parent[root]
	}
	// 뿌리까지 올라가는데 중간에 만나는 노드들을 이미 알고있는 뿌리로 다이렉트로 연결
	for trail = node; trail != root; trail = front {
		front = parent[trail]
		parent[trail] = root
	}
	return root
}
