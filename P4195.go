// Union-find
// https://www.acmicpc.net/problem/4195

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

var parent []int = make([]int, 200005)

func Fn(F int) {
	for i := 0; i < 200005; i++ {
		parent[i] = -1
	}

	nameMap := make(map[string]int)
	cnt := 1

	for i := 0; i < F; i++ {
		N1, N2 := scanString(), scanString()
		if nameMap[N1] == 0 {
			nameMap[N1] = cnt
			cnt++
		}
		if nameMap[N2] == 0 {
			nameMap[N2] = cnt
			cnt++
		}

		union(nameMap[N1], nameMap[N2])
		fmt.Fprintln(bw, abs(parent[find(nameMap[N1])]))
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Tc := scanInt()
	for i := 0; i < Tc; i++ {
		Fn(scanInt())
	}
}

// weighting Rule
func union(x_, y_ int) {
	x := find(x_)
	y := find(y_)

	// root가 같다면 사이클 형성
	if x == y {
		return
	}

	// 적은 node를 leaf로 가진 node를 아래에 붙임
	if parent[x] < parent[y] {
		parent[x] = parent[x] + parent[y]
		parent[y] = x
	} else {
		parent[y] = parent[y] + parent[x]
		parent[x] = y
	}
}

// Collapsing Rule
func find(node int) int {
	var root, trail, front int
	// node(left)에서 출발해서 parent값이 음수인(뿌리)노드를 탐색
	for root = node; parent[root] >= 0; {
		root = parent[root]
	}
	// 뿌리까지 올라가는데 중간에 만나는 노드들을 이미 알고있는 뿌리로 다이렉트로 연결
	for trail = node; trail != root; trail = front {
		front = parent[trail]
		parent[trail] = root
	}
	return root
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
