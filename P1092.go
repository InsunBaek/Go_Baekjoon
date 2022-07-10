// Greedy
// https://www.acmicpc.net/problem/1092

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func Fn() {

	N := scanInt()
	cranes := make([]int, N)
	for i := 0; i < N; i++ {
		cranes[i] = scanInt()
	}

	M := scanInt()
	boxes := make([]int, M)

	for i := 0; i < M; i++ {
		boxes[i] = scanInt()
	}

	sort.Slice(cranes, func(i, j int) bool {
		return cranes[i] > cranes[j]
	})
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i] > boxes[j]
	})

	if cranes[0] < boxes[0] {
		fmt.Fprintln(bw, "-1")
	} else {
		time := 0
		for len(boxes) != 0 {
			idx := 0
			for i := 0; i < len(cranes); {
				if idx == len(boxes) {
					break
				} else if cranes[i] >= boxes[idx] {
					boxes = append(boxes[:idx], boxes[idx+1:]...)
					i++
				} else {
					idx++
				}
			}
			time++
		}
		fmt.Fprintln(bw, time)
	}
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
