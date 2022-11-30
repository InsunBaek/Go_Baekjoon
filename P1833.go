// MST
// https://www.acmicpc.net/problem/1833

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

var constCost [201][201]int

func Fn() {
	pq := &EdgeHeap{}
	heap.Init(pq)

	N := scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			cost := scanInt()
			constCost[i][j] = cost
		}
	}

	sumCost, cnt := 0, 0
	UFInit(N + 1)

	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			cost := constCost[i][j]
			if cost < 0 {
				sumCost += (-cost)
				union(i, j)
			} else {
				heap.Push(pq, &Edge{V1: i, V2: j, Cost: cost})
			}
		}
	}

	edgeSlice := make([]Edge, 0)
	for pq.Len() > 0 {
		curEdge := heap.Pop(pq).(Edge)
		vtx1, vtx2, cost := curEdge.V1, curEdge.V2, curEdge.Cost
		if find(vtx1) == find(vtx2) {
			continue
		} else {
			union(vtx1, vtx2)
			sumCost += cost
			cnt++
			edgeSlice = append(edgeSlice, Edge{V1: vtx1, V2: vtx2, Cost: cost})
		}
	}

	fmt.Fprintln(bw, sumCost, cnt)
	for i := 0; i < len(edgeSlice); i++ {
		fmt.Fprintln(bw, edgeSlice[i].V1, edgeSlice[i].V2)
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

/* Priority Queue */
type Edge struct {
	V1, V2, Cost int
}

type EdgeHeap []*Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(v interface{}) { *h = append(*h, v.(*Edge)) }
func (h *EdgeHeap) Pop() interface{} {
	old := *h
	leng := len(old)
	x := old[leng-1]
	*h = old[0 : leng-1]
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
