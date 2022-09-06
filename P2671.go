// Regexp
// https://www.acmicpc.net/problem/2671

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	exp, _ := regexp.Compile("^(100+1+|01)+$")

	sc.Scan()

	if exp.MatchString(sc.Text()) {
		fmt.Println("SUBMARINE")
	} else {
		fmt.Println("NOISE")
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
