// Map, Sorting
// https://www.acmicpc.net/problem/16165

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func Fn() {

	info_group := make(map[string]([]string))
	info_mem := make(map[string]string)

	N, M := scanInt(), scanInt()

	for i := 0; i < N; i++ {
		groupName := scanString()
		member := scanInt()
		info_group[groupName] = make([]string, member)

		for j := 0; j < member; j++ {
			name := scanString()
			info_group[groupName][j] = name
			info_mem[name] = groupName
		}
	}

	for pro := 0; pro < M; pro++ {
		Name := scanString()
		num := scanInt()

		if num == 1 {
			fmt.Fprintln(bw, info_mem[Name])
		} else {
			sort.Slice(info_group[Name], func(i, j int) bool {
				return info_group[Name][i] < info_group[Name][j]
			})
			for i := 0; i < len(info_group[Name]); i++ {
				fmt.Fprintln(bw, info_group[Name][i])
			}
		}
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
