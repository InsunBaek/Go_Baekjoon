package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)
var graph []Edge

type Edge struct {
	toList []int
}

func Fn() bool {

	V, E := scanInt(), scanInt()
	graph = make([]Edge, V+1)
	color := make([]string, V+1)
	visited := make([]bool, V+1)

	for i := 0; i <= V; i++ {
		graph[i].toList = make([]int, 0)
	}

	for i := 0; i < E; i++ {
		v1, v2 := scanInt(), scanInt()
		graph[v1].toList = append(graph[v1].toList, v2)
		graph[v2].toList = append(graph[v2].toList, v1)
	}

	for i := 1; i <= V; i++ {
		queue := NewQueue()
		if visited[i] == false {
			queue.Push(i)
			color[i] = "black"
			visited[i] = true
		} else {
			continue
		}

		for !queue.Empty() {
			now := queue.Pop().(int)

			for i := 0; i < len(graph[now].toList); i++ {
				to := graph[now].toList[i]
				if visited[to] {
					if color[now] == "white" {
						if color[to] == "white" {
							return false
						}
					} else if color[now] == "black" {
						if color[to] == "black" {
							return false
						}
					}
				} else {
					if color[now] == "black" {
						color[to] = "white"
						visited[to] = true
					} else {
						color[to] = "black"
						visited[to] = true
					}
					queue.Push(to)
				}
			}
		}
	}
	return true
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Tc := scanInt()
	for Tc > 0 {
		fact := Fn()
		if fact {
			fmt.Fprintln(bw, "YES")
		} else {
			fmt.Fprintln(bw, "NO")
		}
		Tc--
	}
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
