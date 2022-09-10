// Math
// https://www.acmicpc.net/problem/2455

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
	people, max := 0, 0

	for i := 0; i < 4; i++ {
		out, in := scanInt(), scanInt()
		people = people - out + in
		if people > max {
			max = people
		}
	}
	fmt.Println(max)
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
