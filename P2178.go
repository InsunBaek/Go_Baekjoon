// Graph, BFS
// https://www.acmicpc.net/problem/2178

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
var dy = [...]int{1, -1, 0, 0}
var dx = [...]int{0, 0, 1, -1}
var graph [][]int
var visit [][]bool
var row, col int
var queue *Queue

type Node struct {
	y   int
	x   int
	cnt int
}
type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front == nil {
		return nil
	}
	return q.v.Remove(front)
}

func (q *Queue) Empty() bool {
	if q.v.Len() == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return q.v.Len()
}

func BFS(y_ int, x_ int) int {
	queue = NewQueue()
	queue.Push(Node{y: y_, x: x_, cnt: 1})
	visit[y_][x_] = true

	for !queue.Empty() {
		node := queue.Pop().(Node)
		y := node.y
		x := node.x

		if y == (row-1) && x == (col-1) {
			return node.cnt
		}

		for i := 0; i < 4; i++ {
			ny := y + dy[i]
			nx := x + dx[i]

			if ny < 0 || nx < 0 || ny >= row || nx >= col {
				continue
			}
			if visit[ny][nx] {
				continue
			}
			if graph[ny][nx] == 1 {
				visit[ny][nx] = true
				queue.Push(Node{y: ny, x: nx, cnt: node.cnt + 1})
			}
		}
	}

	return -1
}

func main() {
	defer bw.Flush()

	br.Scan()
	line := strings.Split(br.Text(), " ")

	row, _ = strconv.Atoi(line[0])
	col, _ = strconv.Atoi(line[1])

	graph = make([][]int, row)
	visit = make([][]bool, row)
	for i := 0; i < row; i++ {
		graph[i] = make([]int, col)
		visit[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		br.Scan()
		token := strings.Split(br.Text(), "")
		for j := 0; j < col; j++ {
			val, _ := strconv.Atoi(token[j])
			graph[i][j] = val
		}
	}

	fmt.Println(BFS(0, 0))
}
