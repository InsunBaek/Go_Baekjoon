// https://www.acmicpc.net/problem/1753
// Dijkstra

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
var INF = math.MaxInt32
var V, E int

type EdgeHeap []*Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(u interface{}) { *h = append(*h, u.(*Edge)) }
func (h *EdgeHeap) Pop() interface{} {
	old := *h
	long := len(old)
	x := old[long-1]
	*h = old[0 : long-1]
	return *x
}

type Edge struct {
	vertex int
	cost   int
}

func Dijkstra(graph map[int][]Edge, start int) []int {

	visited := make([]bool, V+1)
	distance := make([]int, V+1)
	for i := 0; i < len(distance); i++ {
		distance[i] = INF
	}

	distance[start] = 0

	pq := &EdgeHeap{}
	heap.Init(pq)
	heap.Push(pq, &Edge{start, 0}) // 시작 정점 삽입

	for pq.Len() > 0 {
		curEdge := heap.Pop(pq).(Edge)
		curVtx := curEdge.vertex
		if visited[curVtx] { // 이미 방문했다면 무시
			continue
		}
		visited[curVtx] = true // 방문 완료

		for i := 0; i < len(graph[curVtx]); i++ { // 현재 정점 기준 연결된 엣지 모두 탐색
			nodeEdge := graph[curVtx][i]
			nextVtx, nextCost := nodeEdge.vertex, nodeEdge.cost
			newCost := distance[curVtx] + nextCost

			// 새로운 최단거리 발견시 업데이트
			if distance[nextVtx] >= newCost {
				distance[nextVtx] = newCost
				heap.Push(pq, &Edge{nextVtx, distance[nextVtx]})
			}
		}
	}

	return distance
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	V, E = scanInt(), scanInt()
	start := scanInt()

	Graph := make(map[int][]Edge, V+1)
	for i := 0; i < E; i++ {
		from, to, cost := scanInt(), scanInt(), scanInt()
		Graph[from] = append(Graph[from], Edge{to, cost})
	}

	// 다익스트라 실행후 start로 부터의 각정점의 최단거리 반환
	distance := Dijkstra(Graph, start)

	for i := 1; i <= V; i++ {
		if distance[i] == math.MaxInt32 {
			fmt.Fprintln(bw, "INF")
		} else {
			fmt.Fprintln(bw, distance[i])
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
