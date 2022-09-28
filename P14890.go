// https://www.acmicpc.net/problem/14890
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
var N, L int

// usage 0 -> 평지, 1 -> 내리막길, 2 -> 오르막길
func RoadAvailable(Idx int, usage, road []int, flag bool) bool {
	if flag {
		if Idx+L >= N {
			return false
		}
		backBlock := road[Idx+L]
		if road[Idx]-backBlock != 1 {
			return false
		}
		for i := Idx; i <= Idx+L; i++ {
			usage[i] = 1
		}
	} else {
		if Idx-L < 0 {
			return false
		}
		frontBlock := road[Idx-L]
		for i := Idx - L; i < Idx; i++ {
			if usage[i] != 0 {
				return false
			}
		}
		if road[Idx]-frontBlock != 1 {
			return false
		}
		for i := Idx - L; i < Idx; i++ {
			usage[i] = 2
		}
	}
	return true
}

func Fn(route []int) bool {
	leng := len(route)
	usage := make([]int, leng)

	preArea := route[0]
	for i := 1; i < len(route); i++ {
		if preArea == route[i] {
			continue
		}

		// 내리막길인 경우
		if preArea > route[i] {
			if !RoadAvailable(i-1, usage, route, true) {
				return false
			}
		}
		// 오르막길인 경우
		if preArea < route[i] {
			if !RoadAvailable(i, usage, route, false) {
				return false
			}
		}
		preArea = route[i]
	}
	return true
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, L = scanInt(), scanInt()

	road := make([][]int, N)
	for i := 0; i < N; i++ {
		road[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			road[i][j] = scanInt()
		}
	}

	myRoad := make([]int, 0)
	cnt := 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			myRoad = append(myRoad, road[i][j])
		}
		if Fn(myRoad) {
			cnt++
		}
		myRoad = make([]int, 0)
	}

	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {
			myRoad = append(myRoad, road[i][j])
		}
		if Fn(myRoad) {
			cnt++
		}
		myRoad = make([]int, 0)
	}

	fmt.Println(cnt)
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
