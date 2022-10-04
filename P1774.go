// MST
// https://www.acmicpc.net/problem/1774

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

var graph []map[int]int

type Star struct {
	X, Y int
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	pq := &EdgeHeap{}
	heap.Init(pq)

	N, M := scanInt(), scanInt()
	UFInit(N + 1)

	starSlice := make([]Star, 0)
	for i := 0; i < N; i++ {
		x, y := scanInt(), scanInt()
		starSlice = append(starSlice, Star{X: x, Y: y})
	}

	for i := 0; i < M; i++ {
		se1, se2 := scanInt()-1, scanInt()-1
		union(se1, se2)
	}

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if find(i) == find(j) {
				continue
			}
			star_1, star_2 := starSlice[i], starSlice[j]
			s1x, s2x := float64(star_1.X), float64(star_2.X)
			s1y, s2y := float64(star_1.Y), float64(star_2.Y)
			distance := math.Sqrt((float64(math.Pow((s1x-s2x), 2)) + math.Pow((s1y-s2y), 2)))

			heap.Push(pq, &Edge{V1: i, V2: j, Cost: distance})
		}
	}

	totalCost := 0.0

	for pq.Len() > 0 {
		myEdge := heap.Pop(pq).(Edge)
		v1, v2 := myEdge.V1, myEdge.V2

		if find(v1) == find(v2) {
			continue
		}

		totalCost += myEdge.Cost
		union(v1, v2)
	}

	fmt.Printf("%.2f", totalCost)
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

/* Priority Queue */
type Edge struct {
	V1, V2 int
	Cost   float64
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
