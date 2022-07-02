// Sorting
// https://www.acmicpc.net/problem/11728

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()

	seq_1 := make([]int, N+1)
	seq_2 := make([]int, M+1)
	seq_3 := make([]int, N+M)

	seq_1[N] = math.MaxInt32
	seq_2[M] = math.MaxInt32

	for i := 0; i < N; i++ {
		seq_1[i] = scanInt()
	}
	for i := 0; i < M; i++ {
		seq_2[i] = scanInt()
	}

	ptr1, ptr2 := 0, 0

	for i := 0; i < N+M; i++ {
		if seq_1[ptr1] <= seq_2[ptr2] {
			seq_3[i] = seq_1[ptr1]
			ptr1++
		} else {
			seq_3[i] = seq_2[ptr2]
			ptr2++
		}
	}

	for i := 0; i < len(seq_3); i++ {
		fmt.Fprint(bw, seq_3[i], " ")
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
