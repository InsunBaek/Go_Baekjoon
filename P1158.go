// DataStructure
// https://www.acmicpc.net/problem/1158

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	N, K := scanInt(), scanInt()

	save := make([]int, 0)
	visited := make([]int, N+1)
	cnt, cur := 0, K

	for {
		if visited[cur] == 0 {
			save = append(save, cur)
			visited[cur] = 1
			cnt++
		}

		if cnt == N {
			break
		}

		inCnt := 0
		for {
			cur = cur + 1
			if cur > N {
				cur = cur - N
			}
			if visited[cur] == 0 {
				inCnt++
			} else {
				continue
			}

			if inCnt == K {
				break
			}
		}
	}

	fmt.Fprint(bw, "<")
	for i := 0; i < len(save)-1; i++ {
		fmt.Fprint(bw, save[i], ", ")
	}
	fmt.Fprint(bw, save[len(save)-1], ">")
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
