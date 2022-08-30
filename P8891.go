// https://www.acmicpc.net/problem/8891
// BruteForce

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func calc(v int) (int, int) {

	x, y, a := 0, 1, 0

	for v > a {
		v = v - a
		a++
		y++
	}

	for i := 0; i < v; i++ {
		x = x + 1
		y = y - 1
	}

	return x, y
}

func revCalc(x, y int) int {

	nX, nY, v := 1, 1, 1

	for {
		if nX == x && nY == y {
			return v
		}

		if nY-1 <= 0 {
			nY = nX + 1
			nX = 1
			v++
		} else {
			nY--
			nX++
			v++
		}
	}
}

func Fn() {

	T := scanInt()

	for i := 0; i < T; i++ {
		a, b := scanInt(), scanInt()
		x1, y1 := calc(a)
		x2, y2 := calc(b)

		fmt.Println(revCalc(x1+x2, y1+y2))
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
