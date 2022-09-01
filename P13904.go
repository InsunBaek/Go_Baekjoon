// Greedy, Sorting
// https://www.acmicpc.net/problem/13904

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

type Work struct {
	dday  int
	point int
}

func Fn() {

	N := scanInt()

	WorkSlice := make([]Work, 1005)
	schedule := make([]int, 1005)

	for i := 0; i < N; i++ {
		d, w := scanInt(), scanInt()
		WorkSlice = append(WorkSlice, Work{dday: d, point: w})
	}

	sort.Slice(WorkSlice, func(i, j int) bool {
		return WorkSlice[i].point > WorkSlice[j].point
	})

	for i := 0; i < N; i++ {
		curWork := WorkSlice[i]
		curDday := curWork.dday
		curPoint := curWork.point

		if schedule[curDday] == 0 {
			schedule[curDday] = curPoint
		} else {
			for {
				if curDday == 1 {
					break
				}

				curDday--
				if schedule[curDday] == 0 {
					schedule[curDday] = curPoint
					break
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(schedule); i++ {
		sum += schedule[i]
	}
	fmt.Println(sum)
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
