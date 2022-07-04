// https://www.acmicpc.net/problem/25175

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, M, K := scanInt(), scanInt(), scanInt()

	if K >= 3 {
		temp := K - 3
		mod := temp % N
		if M+mod > N {
			temp := M + mod - N
			fmt.Fprintln(bw, temp)
		} else {
			fmt.Fprintln(bw, (M + mod))
		}
	} else {
		temp := 3 - (K)
		mod := temp % N
		if M-mod <= 0 {
			fmt.Fprintln(bw, (N + (M - mod)))
		} else {
			fmt.Fprintln(bw, (M - mod))
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
