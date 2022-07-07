// Dijkstra
// https://www.acmicpc.net/problem/10473

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)
var Graph = [102]Vertex{}
var INF = math.MaxFloat64
var cannons int

type Vertex struct {
	row  float64
	col  float64
	time float64
}

func twoPointDistance(fromRow float64, fromCol float64, toRow float64, toCol float64) float64 {
	return math.Sqrt(math.Pow((toRow-fromRow), 2) + math.Pow((toCol-fromCol), 2))
}

func calcTime(fromRow float64, fromCol float64, toRow float64, toCol float64, Cannon bool) float64 {
	if Cannon {
		distance := twoPointDistance(fromRow, fromCol, toRow, toCol)
		if distance >= 50 {
			return (distance-50)/5 + 2
		} else {
			return (50-distance)/5 + 2
		}
	} else {
		return twoPointDistance(fromRow, fromCol, toRow, toCol) / 5
	}
}

func main() {
	defer bw.Flush()

	sRow, sCol := scanVertex()
	eRow, eCol := scanVertex()

	Graph[0] = Vertex{row: sRow, col: sCol}
	Graph[1] = Vertex{row: eRow, col: eCol}

	cannons = scanInt()
	for i := 0; i < cannons; i++ {
		cannonRow, cannonCol := scanVertex()
		Graph[i+2] = Vertex{row: cannonRow, col: cannonCol}
	}

	distance := make([]float64, cannons+2)
	for i := 1; i < len(distance); i++ {
		distance[i] = INF
	}

	time := Dijkstra(Graph[:], distance)
	fmt.Printf("%5f", time)
}

func Dijkstra(graph []Vertex, distance []float64) float64 {
	pq := &VertexHeap{}
	heap.Init(pq)
	heap.Push(pq, &Vertex{row: Graph[0].row, col: Graph[0].col, time: 0})
	usedCannon := false

	for pq.Len() > 0 {

		curVertex := heap.Pop(pq).(Vertex)
		curRow, curCol, curTime := curVertex.row, curVertex.col, curVertex.time

		for i := 1; i < cannons+2; i++ {
			nextVertex := Graph[i]
			nextRow, nextCol := nextVertex.row, nextVertex.col
			walkTime := calcTime(curRow, curCol, nextRow, nextCol, false)
			cannonTime := calcTime(curRow, curCol, nextRow, nextCol, true)

			minTime := math.Min(walkTime, cannonTime)
			if !usedCannon {
				minTime = walkTime
			}
			if distance[i] > (curTime + minTime) {
				distance[i] = curTime + minTime
				heap.Push(pq, &Vertex{nextRow, nextCol, distance[i]})
			}
		}
		usedCannon = true
	}
	return distance[1]
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func scanVertex() (float64, float64) {
	sc.Scan()
	tokens := strings.Fields(sc.Text())
	r, _ := strconv.ParseFloat(tokens[0], 64)
	c, _ := strconv.ParseFloat(tokens[1], 64)

	return r, c
}

type VertexHeap []*Vertex

func (h VertexHeap) Len() int           { return len(h) }
func (h VertexHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h VertexHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *VertexHeap) Push(u interface{}) { *h = append(*h, u.(*Vertex)) }
func (h *VertexHeap) Pop() interface{} {
	old := *h
	long := len(old)
	x := old[long-1]
	*h = old[0 : long-1]
	return *x
}
