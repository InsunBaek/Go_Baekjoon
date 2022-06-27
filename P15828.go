// Queue
// https://www.acmicpc.net/problem/15828

package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

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

func main() {
	defer bw.Flush()

	queue := *NewQueue()

	br.Scan()
	N, _ := strconv.Atoi(br.Text())

	for {
		br.Scan()
		Info, _ := strconv.Atoi(br.Text())
		if Info == -1 {
			break
		}

		if Info == 0 {
			queue.Pop()
		} else {
			if queue.Size() != N {
				queue.Push(Info)
			}
		}
	}

	for !queue.Empty() {
		bw.WriteString(strconv.Itoa(queue.Pop().(int)) + " ")
	}
}
