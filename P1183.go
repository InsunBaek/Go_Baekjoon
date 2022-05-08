// Math, Sorting
// https://www.acmicpc.net/problem/1183

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {

	br.Scan()
	N, _ := strconv.Atoi(br.Text())

	ints := make([]int, N)
	for i := 0; i < N; i++ {
		br.Scan()
		slice := strings.Split(br.Text(), " ")
		A, _ := strconv.Atoi(slice[0])
		B, _ := strconv.Atoi(slice[1])

		ints[i] = A - B
	}

	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})

	if N%2 == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(ints[N/2] - ints[(N/2)-1] + 1)
	}
}
