// Implementation
// https://www.acmicpc.net/submit/3961

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var rd = bufio.NewReader(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

type Item struct {
	str string
	cnt int
}

func OutputIndex(ch string) (int, int) {

	row, col := 0, 0
	switch ch {
	case "q":
		row, col = 0, 0
	case "w":
		row, col = 0, 1
	case "e":
		row, col = 0, 2
	case "r":
		row, col = 0, 3
	case "t":
		row, col = 0, 4
	case "y":
		row, col = 0, 5
	case "u":
		row, col = 0, 6
	case "i":
		row, col = 0, 7
	case "o":
		row, col = 0, 8
	case "p":
		row, col = 0, 9
	case "a":
		row, col = 1, 0
	case "s":
		row, col = 1, 1
	case "d":
		row, col = 1, 2
	case "f":
		row, col = 1, 3
	case "g":
		row, col = 1, 4
	case "h":
		row, col = 1, 5
	case "j":
		row, col = 1, 6
	case "k":
		row, col = 1, 7
	case "l":
		row, col = 1, 8
	case "z":
		row, col = 2, 0
	case "x":
		row, col = 2, 1
	case "c":
		row, col = 2, 2
	case "v":
		row, col = 2, 3
	case "b":
		row, col = 2, 4
	case "n":
		row, col = 2, 5
	case "m":
		row, col = 2, 6
	}

	return row, col
}

func Abs(a int, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

func Fn(stdString string, tempString string) int {

	value := 0
	for i := 0; i < len(stdString); i++ {
		ch_1 := string(stdString[i])
		ch_2 := string(tempString[i])

		row_1, col_1 := OutputIndex(ch_1)
		row_2, col_2 := OutputIndex(ch_2)

		value += Abs(row_1, row_2) + Abs(col_1, col_2)
	}
	return value
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	Tc := scanInt()

	for i := 0; i < Tc; i++ {

		sc.Scan()
		stdString := sc.Text()

		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())

		slice := make([]Item, 0)

		for j := 0; j < num; j++ {
			sc.Scan()
			word := sc.Text()
			slice = append(slice, Item{str: word, cnt: Fn(stdString, word)})
		}

		sort.Slice(slice, func(i, j int) bool {
			if slice[i].cnt != slice[j].cnt {
				return slice[i].cnt < slice[j].cnt
			} else {
				return slice[i].str < slice[j].str
			}
		})

		for k := 0; k < len(slice); k++ {
			fmt.Fprintln(bw, slice[k].str, slice[k].cnt)
		}
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
