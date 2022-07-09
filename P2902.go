// String
// https://www.acmicpc.net/problem/2902

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func Fn() {

	sc.Scan()
	name := []rune(sc.Text())

	front := name[0]
	answer := ""

	for i := 0; i < len(name); i++ {
		if name[i] == '-' {
			answer += string(front)
			front = name[i+1]
		}
	}
	answer += string(front)
	fmt.Fprintln(bw, answer)
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
