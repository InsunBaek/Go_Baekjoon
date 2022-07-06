// Dijkstra
// https://www.acmicpc.net/problem/4485

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)
var N int
var INF = math.MaxInt32
var square [][]int
var dy = [...]int{1, -1, 0, 0}
var dx = [...]int{0, 0, 1, -1}

/* Priority Queue */
type CostHeap []*Cost

func (h CostHeap) Len() int           { return len(h) }
func (h CostHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h CostHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CostHeap) Push(u interface{}) { *h = append(*h, u.(*Cost)) }
func (h *CostHeap) Pop() interface{} {
	old := *h
	long := len(old)
	x := old[long-1]
	*h = old[0 : long-1]
	return *x
}

/* Graph */
type Cost struct {
	row  int
	col  int
	cost int
}

type Edge struct {
	toList []Cost
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	cnt := 0

	for {
		N = scanInt()
		if N == 0 {
			break
		}

		square = make([][]int, N)
		for i := 0; i < N; i++ {
			square[i] = make([]int, N)
		}

		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				blackRupee := scanInt()
				square[i][j] = blackRupee
			}
		}

		visited := make([][]bool, N)
		distance := make([][]int, N)
		for i := 0; i < N; i++ {
			visited[i] = make([]bool, N)
			distance[i] = make([]int, N)
		}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				distance[i][j] = INF
			}
		}
		cnt += 1
		wentDistance := Dijkstra(square, visited, distance, 0, 0)
		fmt.Fprint(bw, "Problem ", cnt)
		fmt.Fprintln(bw, ":", wentDistance)
	}
}

func canVisit(N int, row int, col int) bool {
	if row < 0 || col < 0 || row >= N || col >= N {
		return false
	}
	return true
}

func Dijkstra(Square [][]int, visited [][]bool, distance [][]int, startRow int, startCol int) int {
	distance[startRow][startCol] = Square[startRow][startCol]

	pq := &CostHeap{}
	heap.Init(pq)
	heap.Push(pq, &Cost{row: startRow, col: startCol, cost: Square[startRow][startCol]})

	for pq.Len() != 0 {
		currentCost := heap.Pop(pq).(Cost)
		minRow := currentCost.row
		minCol := currentCost.col

		if visited[minRow][minCol] {
			continue
		}
		visited[minRow][minCol] = true

		for i := 0; i < 4; i++ {
			nRow := minRow + dy[i]
			nCol := minCol + dx[i]

			if canVisit(N, nRow, nCol) {
				cost := square[nRow][nCol]
				if !visited[nRow][nCol] && distance[nRow][nCol] > distance[minRow][minCol]+cost {
					distance[nRow][nCol] = distance[minRow][minCol] + cost
					heap.Push(pq, &Cost{row: nRow, col: nCol, cost: distance[nRow][nCol]})
				}
			}
		}
	}

	return distance[N-1][N-1]
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
