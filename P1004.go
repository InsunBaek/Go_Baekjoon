// Geometry
// https://www.acmicpc.net/problem/1004

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

func Distance(sY int, sX int, Y int, X int) float64 {
	return math.Sqrt(math.Pow(float64(sY-Y), 2) + math.Pow(float64(sX-X), 2))
}

func Fn() int {
	cnt := 0

	X1, Y1 := scanInt(), scanInt()
	X2, Y2 := scanInt(), scanInt()
	n := scanInt()

	for i := 0; i < n; i++ {
		nX := scanInt()
		nY := scanInt()
		r := float64(scanInt())

		var flag_1 bool
		var flag_2 bool
		if Distance(Y1, X1, nY, nX) < r {
			flag_1 = true
		}
		if Distance(Y2, X2, nY, nX) < r {
			flag_2 = true
		}

		if flag_1 && flag_2 {
			continue
		} else if flag_1 {
			cnt++
		} else if flag_2 {
			cnt++
		}
	}
	return cnt
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	tc := scanInt()
	for i := 0; i < tc; i++ {
		fmt.Fprintln(bw, Fn())
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return int(v)
}
