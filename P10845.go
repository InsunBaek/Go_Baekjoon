// Queue
// https://www.acmicpc.net/problem/10845

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Queue struct {
	v []int
}

func (q *Queue) push(value int) {
	q.v = append(q.v, value)
}

func (q *Queue) pop() int {
	if len(q.v) == 0 {
		return -1
	}
	front, vlist := q.v[0], q.v[1:]
	q.v = vlist
	return front
}

func (q Queue) size() int {
	return len(q.v)
}

func (q Queue) empty() int {
	if q.size() == 0 {
		return 1
	} else {
		return 0
	}
}

func (q Queue) front() int {
	if len(q.v) == 0 {
		return -1
	} else {
		return q.v[0]
	}
}

func (q Queue) back() int {
	if len(q.v) == 0 {
		return -1
	} else {
		return q.v[q.size()-1]
	}
}

func main() {
	q := Queue{}

	br.Scan()
	try, _ := strconv.Atoi(br.Text())

	for i := 0; i < try; i++ {
		br.Scan()
		command := br.Text()

		switch command {
		case "pop":
			bw.WriteString(strconv.Itoa(q.pop()) + "\n")
		case "size":
			bw.WriteString(strconv.Itoa(q.size()) + "\n")
		case "empty":
			bw.WriteString(strconv.Itoa(q.empty()) + "\n")
		case "front":
			bw.WriteString(strconv.Itoa(q.front()) + "\n")
		case "back":
			bw.WriteString(strconv.Itoa(q.back()) + "\n")
		default:
			slice := strings.Split(command, " ")
			value, _ := strconv.Atoi(slice[1])
			q.push(value)
		}
	}

	bw.Flush()
}
