// https://www.acmicpc.net/problem/14891
// Simulation

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Topny struct {
	status []int
}

func Rotate(topnyStatus []int, RL bool) []int { // RL = True 이면 반시계회전
	// 반시계 회전
	if RL {
		tempStatus := make([]int, 8)
		tempStatus[7] = topnyStatus[0]
		for i := 0; i < 7; i++ {
			tempStatus[i] = topnyStatus[i+1]
		}
		return tempStatus
	}

	// 시계 회전
	tempStatus := make([]int, 8)
	tempStatus[0] = topnyStatus[7]
	for i := 0; i < 7; i++ {
		tempStatus[i+1] = topnyStatus[i]
	}
	return tempStatus
}

// RL = true 이면 반시계, RL = false 이면 시계
func Fn(Topnys []Topny, num int, RL, first, left, right bool) {
	nowTopny := Topnys[num-1]
	nowTopnyLeft := nowTopny.status[6]
	nowTopnyRight := nowTopny.status[2]

	nextLeft := num - 1
	nextRight := num + 1

	Topnys[num-1].status = Rotate(Topnys[num-1].status, RL)

	if first {
		// 왼쪽 톱니
		if nextLeft > 0 {
			leftTopnyRight := Topnys[nextLeft-1].status[2]
			if leftTopnyRight != nowTopnyLeft {
				// fmt.Println(num, " 에서 ", nextLeft, " 로 진행합니다.", leftTopnyRight, " != ", Topnys[nextLeft-1].status[2])
				Fn(Topnys, nextLeft, !RL, false, true, false)
			}
		}
		// 오른쪽 톱니
		if nextRight <= 4 {
			rightTopnyleft := Topnys[nextRight-1].status[6]
			if rightTopnyleft != nowTopnyRight {
				// fmt.Println(num, " 에서 ", nextRight, " 로 진행합니다.", rightTopnyleft, " != ", Topnys[nextLeft-1].status[2])
				Fn(Topnys, nextRight, !RL, false, false, true)
			}
		}
	} else {
		if left {
			if nextLeft > 0 {
				leftTopnyRight := Topnys[nextLeft-1].status[2]
				if leftTopnyRight != nowTopnyLeft {
					// fmt.Println(num, " 에서 ", nextLeft, " 로 진행합니다.", leftTopnyRight, " != ", Topnys[nextLeft-1].status[2])
					Fn(Topnys, nextLeft, !RL, false, true, false)
				}
			}
		} else if right {
			if nextRight <= 4 {
				rightTopnyleft := Topnys[nextRight-1].status[6]
				if rightTopnyleft != nowTopnyRight {
					// fmt.Println(num, " 에서 ", nextRight, " 로 진행합니다.", rightTopnyleft, " != ", Topnys[nextLeft-1].status[2])
					Fn(Topnys, nextRight, !RL, false, false, true)
				}
			}
		}
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Topnys := make([]Topny, 4)

	for i := 0; i < 4; i++ {
		Topnys[i].status = make([]int, 8)
		status := scanString()
		for j := 0; j < 8; j++ {
			Topnys[i].status[j] = int(status[j] - 48)
		}
	}

	try := scanInt()
	for i := 0; i < try; i++ {
		num, direct := scanInt(), scanInt()

		if direct == -1 {
			Fn(Topnys, num, true, true, true, true)
		} else {
			Fn(Topnys, num, false, true, true, true)
		}
	}

	score := 0
	val := 1

	for i := 0; i < len(Topnys); i++ {
		if Topnys[i].status[0] == 1 {
			score = score + val
		}
		val = val * 2
	}
	fmt.Fprintln(bw, score)
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
