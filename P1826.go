// https://www.acmicpc.net/problem/1826
// Greedy

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type GasStation struct {
	location int
	gas      int
}

type GasHeap []*GasStation

func (h GasHeap) Len() int           { return len(h) }
func (h GasHeap) Less(i, j int) bool { return h[i].gas > h[j].gas }
func (h GasHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GasHeap) Push(u interface{}) { *h = append(*h, u.(*GasStation)) }
func (h *GasHeap) Pop() interface{} {
	old := *h
	long := len(old)
	x := old[long-1]
	*h = old[0 : long-1]
	return *x
}

func Fn() {
	N := scanInt()
	gasSlice := make([]GasStation, 0)

	for i := 0; i < N; i++ {
		a, b := scanInt(), scanInt()
		gasSlice = append(gasSlice, GasStation{location: a, gas: b})
	}
	L, P := scanInt(), scanInt()

	sort.Slice(gasSlice, func(i, j int) bool {
		return gasSlice[i].location < gasSlice[j].location
	})

	pq := &GasHeap{}
	heap.Init(pq)

	idx, curLocation, cnt := 0, 0, 0

	for {

		// 현재 연료 상태로 목적지 L 도달 가능할시 카운트 출력 후 종료
		if curLocation+P >= L {
			fmt.Println(cnt)
			os.Exit(0)
		}

		// 현재 연료 상태로 도달 가능한 모든 주유소를 큐에 삽입
		for {
			if idx < N && curLocation+P >= gasSlice[idx].location {
				heap.Push(pq, &GasStation{gasSlice[idx].location, gasSlice[idx].gas})
				idx++
			} else {
				break
			}
		}

		// 큐가 비어있다면 L에 도달 불가능
		if pq.Len() == 0 {
			fmt.Println("-1")
			break
		}

		// 큐에서 최적의 주유소 하나를 뽑는데 기준은 가장 충전량이 많은 주유소
		nextGas := heap.Pop(pq).(GasStation)
		// 1. 뽑은 주유소가 현재 위치보다 앞에 있는경우 이동후 거리, 연료 계산, cnt++
		// 2. 뽑은 주유소가 현재 위치보다 뒤에 있는경우 연료추가후 cnt++
		if curLocation+P >= nextGas.location {
			P = P - (nextGas.location - curLocation) + nextGas.gas
			curLocation = nextGas.location
			cnt++
		} else {
			P += nextGas.gas
			cnt++
		}
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
