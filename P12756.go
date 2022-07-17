// Math
// https://www.acmicpc.net/problem/12756

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var rd = bufio.NewReader(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	aAttack, aDefen := scanInt(), scanInt()
	bAttack, bDefen := scanInt(), scanInt()

	for {
		aDefen -= bAttack
		bDefen -= aAttack

		if aDefen <= 0 && bDefen <= 0 {
			fmt.Fprintln(bw, "DRAW")
			break
		} else if aDefen <= 0 {
			fmt.Fprintln(bw, "PLAYER B")
			break
		} else if bDefen <= 0 {
			fmt.Fprintln(bw, "PLAYER A")
			break
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
