// Queue
// https://www.acmicpc.net/problem/2164

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Queue struct {
	item chan interface{}
	cnt  int
}

func (q *Queue) Push(val interface{}) {
	q.item <- val
	q.cnt += 1
}

func (q *Queue) Pop() interface{} {
	q.cnt -= 1
	return <-q.item
}

func (q *Queue) Size() int {
	return q.cnt
}

func (q *Queue) Empty() bool {
	if q.cnt == 0 {
		return true
	}
	return false
}

func main() {

	br.Scan()
	N, _ := strconv.Atoi(br.Text())

	queue := Queue{item: make(chan interface{}, N), cnt: 0}

	for i := 1; i <= N; i++ {
		queue.Push(i)
	}

	for {
		card := queue.Pop().(int)
		if queue.Empty() {
			fmt.Println(card)
			break
		}
		queue.Push(queue.Pop())
	}

}
