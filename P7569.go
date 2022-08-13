// https://www.acmicpc.net/problem/7569
// BFS

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

var queue *Queue
var M, N, H int
var box [][][]int
var visit [][][]int
var zeroCnt, oneCnt = 0, 0
var time, queueSize, nextQueueSize = 0, 0, 0

var dy = [...]int{1, -1, 0, 0, 0, 0}
var dx = [...]int{0, 0, 1, -1, 0, 0}
var dz = [...]int{0, 0, 0, 0, 1, -1}

type XYZ struct {
	x, y, z int
}

func makeBoxVisit(M, N, H int) {
	box = make([][][]int, M)
	visit = make([][][]int, M)
	for i := 0; i < len(box); i++ {
		box[i] = make([][]int, N)
		visit[i] = make([][]int, N)
	}
	for i := 0; i < len(box); i++ {
		for j := 0; j < len(box[i]); j++ {
			box[i][j] = make([]int, H)
			visit[i][j] = make([]int, H)
		}
	}
}

func floorInput(h int) {
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			val := scanInt()
			box[i][j][h] = val
			if val == 0 {
				zeroCnt++
			} else if val == 1 {
				queue.Push(XYZ{i, j, h})
				queueSize++
				oneCnt++
			}
		}
	}
}

func Fn() {
	queue = NewQueue()
	N, M, H = scanInt(), scanInt(), scanInt()

	makeBoxVisit(M, N, H)

	for i := 0; i < H; i++ {
		floorInput(i)
	}
	// 익지 않은 토마토가 없을경우 0출력후 종료
	if zeroCnt == 0 {
		fmt.Println("0")
		os.Exit(0)
	}

	ripenTomato, time := 0, 0
	for !queue.Empty() {
		// queueSize -> 현재 시간(cnt)에 익지않은 토마토의 갯수
		for i := 0; i < queueSize; i++ {
			// 익지않은 토마토의 좌표값을 추출
			xyz := queue.Pop().(XYZ)
			nowX, nowY, nowZ := xyz.x, xyz.y, xyz.z

			// 상하좌우, 위, 아래 6공간에 대해서 탐색
			for i := 0; i < 6; i++ {
				newX, newY, newZ := nowX+dx[i], nowY+dy[i], nowZ+dz[i]

				// 박스에서 벗어나면 무시
				if (0 > newX || newX >= M) || (0 > newY || newY >= N) || (0 > newZ || newZ >= H) {
					continue
				}
				// 이미 방문한 좌표라면 무시
				if visit[newX][newY][newZ] == 1 {
					continue
				}
				// 해당 위치에 익지않은 토마토가 있는경우
				if box[newX][newY][newZ] == 0 {
					queue.Push(XYZ{newX, newY, newZ})
					// 큐에 넣고 방문처리
					visit[newX][newY][newZ] = 1
					nextQueueSize++
					ripenTomato++
				}
			}
		}
		time++
		queueSize = nextQueueSize
		nextQueueSize = 0

		// 모든 토마토가 익은경우
		if ripenTomato == zeroCnt {
			fmt.Println(time)
			os.Exit(0)
		}
	}

	// BFS탐색이 끝나도 익지않은 토마토가 남은경우
	fmt.Println("-1")
	os.Exit(0)
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

/* ------- Queue ------- */
type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front == nil {
		return nil
	}
	return q.v.Remove(front)
}

func (q *Queue) Empty() bool {
	if q.v.Len() == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return q.v.Len()
}
