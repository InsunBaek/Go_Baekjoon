// https://www.acmicpc.net/problem/1966
// Queue

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	queue := NewQueue()
	slice := make([]int, 0)

	N, M := scanInt(), scanInt()

	for i := 0; i < N; i++ {
		v := scanInt()
		slice = append(slice, v)
		queue.Push(v)
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})

	nowPtr, maxPtr, order := M, 0, 1

	for {
		peek := queue.Pop()
		if nowPtr == 0 {
			if peek == slice[maxPtr] {
				fmt.Fprintln(bw, order)
				break
			} else {
				queue.Push(peek)
				nowPtr = queue.Size() - 1
			}
			continue
		}

		if peek == slice[maxPtr] {
			order++
			maxPtr++
			nowPtr -= 1
		} else {
			queue.Push(peek)
			nowPtr -= 1
		}
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Tc := scanInt()
	for Tc > 0 {
		Fn()
		Tc--
	}

}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

/* ----- Queue ----- */

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
