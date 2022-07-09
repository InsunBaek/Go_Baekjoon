// Sorting
// https://www.acmicpc.net/problem/2910

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

type Unit struct {
	value     int
	frequency int
	priority  int
}

func Fn() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)
	N, _ := scanInt(), scanInt()

	UnitSlice := make([]Unit, N+1)
	onValue := make([]int, 0)

	m, mm := make(map[int]int), make(map[int]int)
	priority := 1

	for i := 0; i < N; i++ {
		value := scanInt()
		if m[value] == 0 {
			m[value] = priority
			priority++
			onValue = append(onValue, value)
		}
		mm[value]++
	}

	for i := 0; i < len(onValue); i++ {
		v := onValue[i]
		UnitSlice = append(UnitSlice, Unit{value: v, frequency: mm[v], priority: m[v]})
	}

	sort.Slice(UnitSlice, func(i, j int) bool {
		if UnitSlice[i].frequency != UnitSlice[j].frequency {
			return UnitSlice[i].frequency > UnitSlice[j].frequency
		} else {
			return UnitSlice[i].priority < UnitSlice[j].priority
		}
	})

	for i := 0; i < len(UnitSlice); i++ {
		for j := 0; j < UnitSlice[i].frequency; j++ {
			fmt.Fprint(bw, UnitSlice[i].value, " ")
		}
	}
}

func main() {
	defer bw.Flush()
	Fn()
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
