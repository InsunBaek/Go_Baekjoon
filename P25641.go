// String
// https://www.acmicpc.net/problem/25641

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N := scanInt()

	sc.Scan()
	stst := sc.Text()

	sCnt, tCnt := 0, 0
	for i := 0; i < N; i++ {
		if stst[i] == 's' {
			sCnt++
		} else {
			tCnt++
		}
	}

	if sCnt == tCnt {
		fmt.Println(stst)
	} else {
		for i := 0; i < N; i++ {
			if stst[i] == 's' {
				sCnt--
			} else {
				tCnt--
			}
			if sCnt == tCnt {
				fmt.Println(stst[i+1:])
				break
			}
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
