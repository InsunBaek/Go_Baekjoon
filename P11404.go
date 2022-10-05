// floydWarshall
// https://www.acmicpc.net/problem/11404

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

var INF int = 2100000000

func floydWarshall(G [][]int, N int) {

	for waypoint := 1; waypoint <= N; waypoint++ {
		for from := 1; from <= N; from++ {
			for to := 1; to <= N; to++ {
				if G[from][waypoint]+G[waypoint][to] < G[from][to] {
					G[from][to] = G[from][waypoint] + G[waypoint][to]
				}
			}
		}
	}

	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if G[i][j] == INF {
				G[i][j] = 0
			}
		}
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			fmt.Print(G[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()

	graph := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		graph[i] = make([]int, N+1)
	}

	for i := 0; i < N+1; i++ {
		for j := 0; j < N+1; j++ {
			graph[i][j] = INF
			if i == j {
				graph[i][j] = 0
			}
		}
	}

	for i := 0; i < M; i++ {
		from, to, cost := scanInt(), scanInt(), scanInt()
		if graph[from][to] > cost {
			graph[from][to] = cost
		}
	}

	floydWarshall(graph, N)
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
