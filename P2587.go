// Math
// https://www.acmicpc.net/problem/2587

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	slice, sum := make([]int, 0), 0
	for i := 0; i < 5; i++ {
		val := scanInt()
		sum += val
		slice = append(slice, val)
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	fmt.Println(sum / 5)
	fmt.Println(slice[2])
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
