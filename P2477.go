// Geometry
// https://www.acmicpc.net/problem/2477

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func isItLeft(prev int, next int) bool {
	if prev == 0 {
		return true
	}

	switch prev {
	case 1:
		if next == 4 {
			return true
		} else {
			return false
		}
	case 2:
		if next == 3 {
			return true
		} else {
			return false
		}
	case 3:
		if next == 1 {
			return true
		} else {
			return false
		}
	case 4:
		if next == 2 {
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	defer bw.Flush()

	br.Scan()
	K, _ := strconv.Atoi(br.Text())

	rightOccurFlag := false
	maxHeight, maxWidth := 0, 0
	prev, prevValue := 0, 0
	smallRect := 0
	firstValue := 0

	for i := 0; i < 6; i++ {
		br.Scan()
		token := strings.Fields(br.Text())

		direction, _ := strconv.Atoi(token[0])
		length, _ := strconv.Atoi(token[1])
		if i == 0 {
			firstValue = length
		}
		if isItLeft(prev, direction) {
			if direction == 1 || direction == 2 {
				if length > maxWidth {
					maxWidth = length
				}
			}
			if direction == 3 || direction == 4 {
				if length > maxHeight {
					maxHeight = length
				}
			}
		} else {
			smallRect = prevValue * length
			rightOccurFlag = true
		}
		prev = direction
		prevValue = length
	}

	if !rightOccurFlag {
		smallRect = prevValue * firstValue
	}

	volume := maxHeight*maxWidth - smallRect
	bw.WriteString(strconv.Itoa(K * volume))
}
