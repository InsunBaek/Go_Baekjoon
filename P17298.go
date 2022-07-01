// Stack
// https://www.acmicpc.net/problem/17298

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

func (st *Stack) Top() interface{} {
	top := st.v.Back()
	if top == nil {
		return nil
	}
	rtn := st.v.Remove(top)
	st.Push(rtn)
	return rtn
}

func (st *Stack) Empty() bool {
	if st.v.Len() == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	stack := *NewStack()
	N := scanInt()
	seq := make([]int, N)
	ans := make([]int, N)
	for i := 0; i < N; i++ {
		seq[i] = scanInt()
	}

	for i := N - 1; i >= 0; i-- {
		for !stack.Empty() && stack.Top().(int) <= seq[i] {
			stack.Pop()
		}

		if stack.Empty() {
			ans[i] = -1
		} else {
			ans[i] = stack.Top().(int)
		}
		stack.Push(seq[i])
	}

	for i := 0; i < N; i++ {
		fmt.Fprint(bw, ans[i], " ")
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
