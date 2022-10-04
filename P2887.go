// MST
// https://www.acmicpc.net/problem/2887

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Star struct {
	point, starNum int
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	pq := &EdgeHeap{}
	heap.Init(pq)

	V := scanInt()
	UFInit(V + 1)

	x_slice := make([]Star, 0)
	y_slice := make([]Star, 0)
	z_slice := make([]Star, 0)

	for i := 1; i <= V; i++ {
		x_, y_, z_ := scanInt(), scanInt(), scanInt()
		x_slice = append(x_slice, Star{point: x_, starNum: i})
		y_slice = append(y_slice, Star{point: y_, starNum: i})
		z_slice = append(z_slice, Star{point: z_, starNum: i})
	}

	sort.Slice(x_slice, func(i, j int) bool {
		return x_slice[i].point < x_slice[j].point
	})
	sort.Slice(y_slice, func(i, j int) bool {
		return y_slice[i].point < y_slice[j].point
	})
	sort.Slice(z_slice, func(i, j int) bool {
		return z_slice[i].point < z_slice[j].point
	})

	for i := 0; i < V-1; i++ {
		xCost := x_slice[i+1].point - x_slice[i].point
		yCost := y_slice[i+1].point - y_slice[i].point
		zCost := z_slice[i+1].point - z_slice[i].point

		heap.Push(pq, &Edge{V1: x_slice[i].starNum, V2: x_slice[i+1].starNum, Cost: xCost})
		heap.Push(pq, &Edge{V1: y_slice[i].starNum, V2: y_slice[i+1].starNum, Cost: yCost})
		heap.Push(pq, &Edge{V1: z_slice[i].starNum, V2: z_slice[i+1].starNum, Cost: zCost})
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
