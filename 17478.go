// Recursion
// https://www.acmicpc.net/problem/17478

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func makeLine(x int) string {
	str := ""
	for i := 0; i < x; i++ {
		str += "____"
	}
	return str
}
func printLine(x int, y int) {
	line := makeLine(y)

	fmt.Print(line)
	fmt.Println("\"재귀함수가 뭔가요?\"")
	if x == 0 {

		fmt.Print(line)
		fmt.Println("\"재귀함수는 자기 자신을 호출하는 함수라네\"")
		fmt.Print(line)
		fmt.Println("라고 답변하였지.")
		return
	}

	fmt.Print(line)
	fmt.Println("\"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.")
	fmt.Print(line)
	fmt.Println("마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.")
	fmt.Print(line)
	fmt.Println("그의 답은 대부분 옳았다고 하네. 그런데 어느 날, 그 선인에게 한 선비가 찾아와서 물었어.\"")

	printLine(x-1, y+1)

	fmt.Print(line)
	fmt.Println("라고 답변하였지.")
}

func main() {
	defer bw.Flush()

	br.Scan()
	N, _ := strconv.Atoi(br.Text())

	fmt.Println("어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다.")
	printLine(N, 0)
}
