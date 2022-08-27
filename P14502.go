// https://www.acmicpc.net/problem/14502
// Graph, BFS, DFS

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

var dy = [...]int{1, -1, 0, 0}
var dx = [...]int{0, 0, 1, -1}
var N, M, max int
var Facil, visited [][]int

type Wall struct {
	Y int
	X int
}

func Fn() {

	N, M = scanInt(), scanInt()
	Facil = make([][]int, N)
	for i := 0; i < N; i++ {
		Facil[i] = make([]int, M)
	}
	Init()
	wallSlice := make([]Wall, 0)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			v := scanInt()
			if v == 0 {
				wallSlice = append(wallSlice, Wall{Y: i, X: j})
			}
			Facil[i][j] = v
		}
	}

	var wall1, wall2, wall3 Wall

	for i := 0; i < len(wallSlice)-2; i++ {
		wall1 = wallSlice[i]
		for j := i + 1; j < len(wallSlice)-1; j++ {
			wall2 = wallSlice[j]
			for k := j + 1; k < len(wallSlice); k++ {
				wall3 = wallSlice[k]

				facil := copy(Facil)
				Init()
				facil[wall1.Y][wall1.X] = 1
				facil[wall2.Y][wall2.X] = 1
				facil[wall3.Y][wall3.X] = 1

				for row := 0; row < N; row++ {
					for col := 0; col < M; col++ {
						if facil[row][col] == 2 && visited[row][col] == 0 {
							DFS(facil, row, col)
						}
					}
				}

				safeArea := calcSafeArea(facil)
				if max < safeArea {
					max = safeArea
				}

				facil[wall1.Y][wall1.X] = 0
				facil[wall2.Y][wall2.X] = 0
				facil[wall3.Y][wall3.X] = 0
			}
		}
	}

	fmt.Println(max)
}

func DFS(facil [][]int, row, col int) {
	facil[row][col] = 2
	visited[row][col] = 1

	for i := 0; i < 4; i++ {
		nRow := row + dy[i]
		nCol := col + dx[i]

		if nRow < 0 || nRow >= N || nCol < 0 || nCol >= M {
			continue
		}
		if visited[nRow][nCol] == 1 {
			continue
		}
		if facil[nRow][nCol] == 0 {
			DFS(facil, nRow, nCol)
		}
	}
}

func copy(Facil [][]int) [][]int {
	copier := make([][]int, N)
	for i := 0; i < N; i++ {
		copier[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			copier[i][j] = Facil[i][j]
		}
	}
	return copier
}

func calcSafeArea(facil [][]int) int {
	safeArea := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if facil[i][j] == 0 {
				safeArea++
			}
		}
	}
	return safeArea
}

func Init() {
	visited = make([][]int, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]int, M)
	}
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
