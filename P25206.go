// Implementation
// https://www.acmicpc.net/problem/25206

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func calcGrade(grade string) (float64, bool) {
	var returnGrade float64
	var pFlag bool = false

	switch grade {
	case "A+":
		returnGrade = 4.5
	case "A0":
		returnGrade = 4.0
	case "B+":
		returnGrade = 3.5
	case "B0":
		returnGrade = 3.0
	case "C+":
		returnGrade = 2.5
	case "C0":
		returnGrade = 2.0
	case "D+":
		returnGrade = 1.5
	case "D0":
		returnGrade = 1.0
	case "F":
		returnGrade = 0.0
	default:
		pFlag = true
	}
	return returnGrade, pFlag
}

func Fn() {

	cnt := 0.0
	sum := 0.0

	for i := 0; i < 20; i++ {
		sc.Scan()
		tokens := strings.Fields(sc.Text())
		point, _ := strconv.ParseFloat(tokens[1], 64)
		grade, pFlag := calcGrade(tokens[2])

		if pFlag {
			continue
		}
		cnt += point
		sum += point * grade
	}

	fmt.Fprintln(bw, sum/cnt)

}

func main() {
	defer bw.Flush()
	Fn()
}

// func scanInt() int {
// 	sc.Scan()
// 	v, _ := strconv.Atoi(sc.Text())
// 	return v
// }
