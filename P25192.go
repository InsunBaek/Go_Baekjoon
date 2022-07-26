// Map
// https://www.acmicpc.net/problem/25192

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
	N := scanInt()

	log := make(map[string]bool)
	gomgom := 0

	for i := 0; i < N; i++ {
		sc.Scan()
		text := sc.Text()
		if text == "ENTER" {
			log = make(map[string]bool)
		} else {
			if !log[text] {
				gomgom++
				log[text] = true
			}
		}
	}

	fmt.Fprintln(bw, gomgom)
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
