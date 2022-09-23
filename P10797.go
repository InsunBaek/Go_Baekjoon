// https://www.acmicpc.net/problem/10797
// IMP

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

	day := scanInt()
	cnt := 0

	for i := 0; i < 5; i++ {
		if day == scanInt() {
			cnt++
		}
	}
	fmt.Fprintln(bw, cnt)
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
