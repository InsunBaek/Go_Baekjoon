// Sorting, Greedy
// https://www.acmicpc.net/problem/1946

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

type Candidate struct {
	D_rank int
	I_rank int
}

func Fn() {
	N := scanInt()

	cnt := 0
	cddSeq_1 := make([]Candidate, N)
	cddSeq_2 := make([]Candidate, N)

	for i := 0; i < N; i++ {
		d_rank, i_rank := scanInt(), scanInt()
		cddSeq_1[i] = Candidate{D_rank: d_rank, I_rank: i_rank}
		cddSeq_2[i] = Candidate{D_rank: d_rank, I_rank: i_rank}
	}

	sort.Slice(cddSeq_1, func(i, j int) bool {
		return cddSeq_1[i].D_rank < cddSeq_1[j].D_rank
	})
	sort.Slice(cddSeq_2, func(i, j int) bool {
		return cddSeq_2[i].I_rank < cddSeq_2[j].I_rank
	})

	m := make(map[int]int)
	visit := make([]bool, N)
	top := 0

	for i := 0; i < len(cddSeq_2); i++ {
		m[cddSeq_2[i].D_rank] = i
	}

	for i := N - 1; i >= 0; i-- {
		candidata := cddSeq_1[i]
		target := m[candidata.D_rank]
		visit[target] = true
		if target == top {
			cnt++
			for {
				top += 1
				if top == N {
					break
				}
				if visit[top] == false {
					break
				}
			}
		}
	}
	fmt.Fprintln(bw, cnt)
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	tc := scanInt()
	for i := 0; i < tc; i++ {
		Fn()
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
