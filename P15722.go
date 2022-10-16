// Implementation
// https://www.acmicpc.net/problem/15722

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Snail struct {
	X, Y, addr, until, cnt int
	direction              int
}

func Fn(n int) {

	snail := Snail{until: 1, direction: 1}
	time := 0

	if n == 0 {
		fmt.Fprintln(bw, 0, 0)
		return
	}
	for {
		if snail.cnt == 2 {
			snail.cnt = 0
			snail.until++
		}

		switch snail.direction {
		// North
		case 1:
			snail.Y += 1
			snail.addr++
		// East
		case 2:
			snail.X += 1
			snail.addr++
		// South:
		case 3:
			snail.Y -= 1
			snail.addr++
		// West:
		case 4:
			snail.X -= 1
			snail.addr++
		}

		if snail.addr == snail.until {
			if snail.direction == 4 {
				snail.direction = 1
			} else {
				snail.direction++
			}
			snail.addr = 0
			snail.cnt++
		}

		time++
		if time == n {
			fmt.Fprintln(bw, snail.X, snail.Y)
			break
		}
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	n := scanInt()
	Fn(n)
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
