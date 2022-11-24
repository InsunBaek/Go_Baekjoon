// Implementation
// https://www.acmicpc.net/problem/10807

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
	slice := make([]int, 0)
	for i := 0; i < N; i++ {
		slice = append(slice, scanInt())
	}
	target := scanInt()
	vCnt := 0
	for i := 0; i < N; i++ {
		if slice[i] == target {
			vCnt++
		}
	}
	fmt.Println(vCnt)
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
