// https://www.acmicpc.net/problem/1439
// String

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fn() {
	num := scanString()

	zeroCnt := 0
	oneCnt := 0
	now := string(num[0])

	for i := 1; i < len(num); i++ {
		if string(num[i]) == now {
			continue
		}

		if now == "0" {
			zeroCnt++
		} else {
			oneCnt++
		}

		now = string(num[i])
	}
	if now == "0" {
		zeroCnt++
	} else {
		oneCnt++
	}

	if zeroCnt <= oneCnt {
		fmt.Fprint(bw, zeroCnt)
	} else {
		fmt.Fprint(bw, oneCnt)
	}
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
