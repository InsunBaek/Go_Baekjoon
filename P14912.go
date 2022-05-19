// Math, BruteForce
// https://www.acmicpc.net/problem/14912

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func calcDigit(x int, target int) int {
	cnt := 0
	for {
		if x == 0 {
			break
		}

		temp := x % 10
		x /= 10
		if temp == target {
			cnt++
		}
	}
	return cnt
}

func main() {

	br.Scan()
	cnt := 0
	line := strings.Split(br.Text(), " ")
	n, _ := strconv.Atoi(line[0])
	d, _ := strconv.Atoi(line[1])

	for i := 1; i <= n; i++ {
		cnt += calcDigit(i, d)
	}

	bw.WriteString(strconv.Itoa(cnt))
	bw.Flush()

}
