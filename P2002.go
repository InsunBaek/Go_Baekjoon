// https://www.acmicpc.net/problem/2002
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

func Fn() {
	N := scanInt()

	info := make(map[string]int)

	for i := 0; i < N; i++ {
		sc.Scan()
		name := sc.Text()

		info[name] = i
	}

	carSlice := make([]string, 0)

	for i := 0; i < N; i++ {
		sc.Scan()
		name := sc.Text()

		carSlice = append(carSlice, name)
	}

	cnt := 0
	lastNum := info[carSlice[len(carSlice)-1]]

	for i := len(carSlice) - 2; i >= 0; i-- {
		curNum := info[carSlice[i]]
		if curNum < lastNum {
			lastNum = curNum
		} else {
			cnt++
		}
	}

	fmt.Fprintln(bw, cnt)
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
