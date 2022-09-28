// https://www.acmicpc.net/problem/2446
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

func Fn(N int) {

	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			fmt.Fprint(bw, " ")
		}
		for j := 0; j < (2*N)-(2*i)-1; j++ {
			fmt.Fprint(bw, "*")
		}
		fmt.Fprintln(bw)
	}

	for i := 1; i < N; i++ {
		for j := 0; j < N-i-1; j++ {
			fmt.Fprint(bw, " ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Fprint(bw, "*")
		}
		fmt.Fprintln(bw)
	}

}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn(scanInt())
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
