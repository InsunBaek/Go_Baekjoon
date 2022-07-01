// Geometry
// https://www.acmicpc.net/problem/1358

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
var W, H, X, Y, P, rightX, topY, radius int

func Distance(sY int, sX int, Y int, X int) float64 {
	return math.Sqrt(math.Pow(float64(sY-Y), 2) + math.Pow(float64(sX-X), 2))
}

func Fn(x int, y int) bool {
	if (y <= topY && y >= Y) && (x <= rightX && x >= X) {
		return true
	}

	if Distance(y, x, Y+radius, X) <= float64(radius) {
		return true
	}
	if Distance(y, x, Y+radius, rightX) <= float64(radius) {
		return true
	}
	return false
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	cnt := 0
	W, H = scanInt(), scanInt()
	X, Y = scanInt(), scanInt()
	P = scanInt()

	rightX = X + W
	topY = Y + H
	radius = H / 2

	for i := 0; i < P; i++ {
		pY := scanInt()
		pX := scanInt()

		if Fn(pY, pX) {
			cnt++
		}
	}

	fmt.Fprintln(bw, cnt)
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return int(v)
}
