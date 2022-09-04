// Regexp
// https://www.acmicpc.net/problem/5637

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	exp, _ := regexp.Compile("[a-zA-Z-]*")
	max, answer := 0, ""

	for {
		sc.Scan()
		tokens := exp.FindAllString(sc.Text(), 10001)

		for i := 0; i < len(tokens); i++ {
			if tokens[i] == ("E-N-D") {
				fmt.Println(strings.ToLower(answer))
				os.Exit(0)
			}

			length := len(tokens[i])
			if length > max {
				max = length
				answer = tokens[i]
			}
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
