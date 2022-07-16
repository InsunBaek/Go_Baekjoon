// Sorting
// https://www.acmicpc.net/problem/12760

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

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()

	cardList := make([][]int, N)
	for i := 0; i < N; i++ {
		cardList[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		tempSlice := make([]int, M)
		for j := 0; j < M; j++ {
			tempSlice[j] = scanInt()
		}

		sort.Slice(tempSlice, func(i, j int) bool {
			return tempSlice[i] > tempSlice[j]
		})
		for j := 0; j < M; j++ {
			cardList[i][j] = tempSlice[j]
		}
	}

	lstSlice := make([]int, N)
	lstCnt := make([]int, 0)
	lstValue := 0
	for i := 0; i < M; i++ {
		value := 0
		for j := 0; j < N; j++ {
			if value < cardList[j][i] {
				value = cardList[j][i]
				lstCnt = make([]int, 0)
				lstCnt = append(lstCnt, j)
			} else if value == cardList[j][i] {
				lstCnt = append(lstCnt, j)
			}
		}

		for j := 0; j < len(lstCnt); j++ {
			lstSlice[lstCnt[j]] += 1
			if lstValue < lstSlice[lstCnt[j]] {
				lstValue = lstSlice[lstCnt[j]]
			}
		}
	}

	prizeSlice := make([]int, 0)
	for i := 0; i < len(lstSlice); i++ {
		if lstSlice[i] == lstValue {
			prizeSlice = append(prizeSlice, i)
		}
	}

	sort.Slice(prizeSlice, func(i, j int) bool {
		return prizeSlice[i] < prizeSlice[j]
	})
	for i := 0; i < len(prizeSlice); i++ {
		fmt.Fprint(bw, prizeSlice[i]+1, " ")
	}
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
