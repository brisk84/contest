package main

import (
	"fmt"
	"sort"
)

type Proc struct {
	Power     int
	BusyUntil int
}

func main() {

	// in := bufio.NewReader(os.Stdin)
	// n := 0 // procs
	// m := 0 // tasks
	// fmt.Fscan(in, &n, &m)
	// // // power := make([]int, n)
	// procs := make([]Proc, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Fscan(in, &procs[i].Power)
	// }
	// sort.Slice(procs, func(i, j int) bool {
	// 	return procs[i].Power < procs[j].Power
	// })

	// minTime := 0
	// maxTime := 0
	// tasks := make(map[int]int, m)
	// for i := 0; i < m; i++ {
	// 	startTime := 0
	// 	workTime := 0
	// 	fmt.Fscan(in, &startTime, &workTime)
	// 	tasks[startTime] = workTime
	// if i == 0 {
	// 	minTime = starTime
	// }
	// 	if i == m-1 {
	// 		maxTime = startTime
	// 	}
	// }

	m := 1 // tasks
	procs := []Proc{
		{1000000, 0},
	}
	sort.Slice(procs, func(i, j int) bool {
		return procs[i].Power < procs[j].Power
	})

	tasks := make(map[int]int, m)
	tasks[1000000000] = 1000000
	// tasks[2] = 5
	// tasks[3] = 7
	// tasks[4] = 10
	// tasks[5] = 5
	// tasks[6] = 100
	// tasks[9] = 2

	totalPower := 0
	minTime := 1000000000
	maxTime := 1000000000

	for curTime := minTime; curTime <= maxTime; curTime++ {
		if curTask, ok := tasks[curTime]; ok {
			for k := range procs {
				if procs[k].BusyUntil == 0 || procs[k].BusyUntil <= curTime {
					procs[k].BusyUntil = curTime + curTask
					curPower := curTask * procs[k].Power
					totalPower += curPower //curTask * procs[k].Power
					xx := curPower
					_ = xx
					break
				}
			}
		}
	}
	fmt.Println(totalPower)

	// for _, task := range tasks {
	// 	for _, proc := range procs {
	// 		if !proc. {
	// 			proc.Busy = true

	// 		}
	// 	}
	// }

	// fmt.Println(procs)
	// fmt.Println(tasks)

}
