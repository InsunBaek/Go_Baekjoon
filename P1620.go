// Map
// https://www.acmicpc.net/problem/1620

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bw.Flush()

	br.Scan()
	token := strings.Fields(br.Text())

	N, _ := strconv.Atoi(token[0])
	M, _ := strconv.Atoi(token[1])
	strMap := make(map[string]int)
	intMap := make(map[int]string)

	for i := 0; i < N; i++ {
		br.Scan()
		poketmon := br.Text()
		strMap[poketmon] = i + 1
		intMap[i+1] = poketmon
	}

	for i := 0; i < M; i++ {
		br.Scan()
		str := br.Text()
		strKey := str
		intKey, _ := strconv.Atoi(str)

		if strMap[strKey] != 0 {
			bw.WriteString(strconv.Itoa(strMap[strKey]) + "\n")
		}
		if intMap[intKey] != "" {
			bw.WriteString(intMap[intKey] + "\n")
		}
	}
}
