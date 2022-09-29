// https://www.acmicpc.net/problem/11866
// Implementation

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn(N, K int) {

	circle := make([]bool, N+1)
	now, ansCnt := 0, 0

	fmt.Fprint(bw, "<")
	for {
		cnt := 0
		for {
			now += 1
			if now > N {
				now = 1
			}
			if circle[now] {
				continue
			} else {
				cnt++
			}
			if cnt == K {
				break
			}
		}
		circle[now] = true
		ansCnt++

		if ansCnt != N {
			fmt.Fprint(bw, now, ", ")
		} else {
			fmt.Fprint(bw, now, ">")
			break
		}
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn(scanInt(), scanInt())
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
