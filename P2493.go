// Stack
// https://www.acmicpc.net/problem/2493

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Tower struct {
	height int
	tag    int
}

func Fn() {

	stack := *NewStack()
	N := scanInt()

	for i := 0; i < N; i++ {
		h := scanInt()
		for !stack.Empty() {
			if stack.Peek().(Tower).height < h {
				stack.Pop()
			} else {
				fmt.Fprint(bw, stack.Peek().(Tower).tag, " ")
				break
			}
		}
		if stack.Empty() {
			fmt.Fprint(bw, "0 ")
		}
		stack.Push(Tower{height: h, tag: i + 1})
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn()
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

/* Stack */
type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (st *Stack) Push(v interface{}) {
	st.v.PushBack(v)
}

func (st *Stack) Peek() interface{} {
	top := st.v.Back()
	if top == nil {
		return nil
	}
	return st.v.Back().Value
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
