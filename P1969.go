// String, Greedy
// https://www.acmicpc.net/problem/1969

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, M int

func Fn() {

	N, M = scanInt(), scanInt()
	dnaSlice := make([]string, 0)
	for i := 0; i < N; i++ {
		dnaSlice = append(dnaSlice, scanString())
	}

	var DNA bytes.Buffer
	sum := 0

	for idx := 0; idx < M; idx++ {
		charMap := make(map[string]int)
		for s := 0; s < N; s++ {
			charMap[string(dnaSlice[s][idx])] += 1
		}

		max, unit := charMap["T"], "T"
		if max <= charMap["G"] {
			max, unit = charMap["G"], "G"
		}
		if max <= charMap["C"] {
			max, unit = charMap["C"], "C"
		}
		if max <= charMap["A"] {
			max, unit = charMap["A"], "A"
		}

		DNA.WriteString(unit)
		sum += (N - max)
	}

	fmt.Println(DNA.String())
	fmt.Println(sum)

}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Fn()
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
