// Greedy
// https://www.acmicpc.net/problem/25972

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Domino struct {
	// a 좌표, l 높이
	a, l int
}

func Fn() {

	N := scanInt()

	dominos := make([]Domino, 0)
	for i := 0; i < N; i++ {
		a, l := scanInt(), scanInt()
		dominos = append(dominos, Domino{a, l})
	}

	// 좌표기준 정렬
	sort.Slice(dominos, func(i, j int) bool {
		return dominos[i].a < dominos[j].a
	})

	cnt, idx := 1, 0

	for {
	LA:
		curDomino := dominos[idx]
		curA, curL := curDomino.a, curDomino.l
		if idx+1 == N {
			break
		}
		idx++
		for {
			// 다음 도미노를 쓰러뜨릴수 있으면 goto 아니면 카운팅
			if curA+curL >= dominos[idx].a {
				goto LA
			} else {
				cnt++
				break
			}
		}
		if idx >= N {
			break
		}
	}

	fmt.Println(cnt)
}
