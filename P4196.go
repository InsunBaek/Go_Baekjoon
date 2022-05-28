// SCC
// https://www.acmicpc.net/problem/4196
// 해당 문제는 SCC와 관련있지만 역방향 탐색은 할 필요가 없음
// 스택의 마지막 node로부터 도달 가능한 node를 모두 한 번에 쓰러트릴수 있는 도미노로 판단.

package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
var stack Stack
var Graph []Edge
var visited []int

type Edge struct {
	edges []int
}

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (st *Stack) Push(v interface{}) {
	st.v.PushBack(v)
}

func (st *Stack) Pop() interface{} {
	top := st.v.Back()
	if top == nil {
		return nil
	}
	return st.v.Remove(top)
}

func (st *Stack) Empty() bool {
	if st.v.Len() == 0 {
		return true
	} else {
		return false
	}
}

func DFS(x int) {
	if visited[x] == 1 {
		return
	}
	visited[x] = 1
	for i := 0; i < len(Graph[x].edges); i++ {
		temp := Graph[x].edges[i]
		DFS(temp)
	}

	stack.Push(x)
}

func runner() {

	br.Scan()
	line := strings.Split(br.Text(), " ")
	N, _ := strconv.Atoi(line[0])
	M, _ := strconv.Atoi(line[1])

	stack = *NewStack()
	Graph = make([]Edge, N+1)
	visited = make([]int, N+1)
	for i := 0; i < N; i++ {
		Graph[i].edges = make([]int, 0)
	}

	for i := 0; i < M; i++ {
		br.Scan()
		line := strings.Split(br.Text(), " ")
		from, _ := strconv.Atoi(line[0])
		to, _ := strconv.Atoi(line[1])

		Graph[from].edges = append(Graph[from].edges, to)
	}

	for i := 1; i <= N; i++ {
		if visited[i] == 0 {
			DFS(i)
		}
	}

	visited = make([]int, N+1)
	ans := 0
	for {
		if stack.Empty() {
			break
		}

		cur := stack.Pop().(int)
		if visited[cur] == 0 {
			DFS(cur)
			ans++
		}
	}
	bw.WriteString(strconv.Itoa(ans) + "\n")
}

func main() {

	br.Scan()
	try, _ := strconv.Atoi(br.Text())

	for i := 0; i < try; i++ {
		runner()
	}
	bw.Flush()

}
