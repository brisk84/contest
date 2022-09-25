package main

import (
	"fmt"
	"sort"
)

type Task struct {
	StartTime int
	WorkTime  int
}

type Proc struct {
	Power     int
	BusyUntil int
}

func main() {

	// in := bufio.NewReader(os.Stdin)

	// n := 4 // procs
	// m := 7 // tasks
	// fmt.Fscan(in, &n, &m)
	// // power := make([]int, n)
	// procs := make([]Proc, n)
	// for i := 0; i < n; i++ {
	// 	// fmt.Fscan(in, &power[i])
	// 	fmt.Fscan(in, &procs[i].Power)
	// }
	// sort.Slice(procs, func(i, j int) bool {
	// 	return procs[i].Power < procs[j].Power
	// })

	// tasks := make([]Task, m)
	// for i := 0; i < m; i++ {
	// 	fmt.Fscan(in, &tasks[i].StartTime)
	// 	fmt.Fscan(in, &tasks[i].WorkTime)
	// }

	// n := 4 // procs
	m := 7 // tasks

	procs := []Proc{
		{3, 0},
		{2, 0},
		{6, 0},
		{4, 0},
	}
	sort.Slice(procs, func(i, j int) bool {
		return procs[i].Power < procs[j].Power
	})

	// tasks := []Task{
	// 	{1, 3},
	// 	{2, 5},
	// 	{3, 7},
	// 	{4, 10},
	// 	{5, 5},
	// 	{6, 100},
	// 	{9, 2},
	// }
	// tasks := make([]int, m)
	tasks := make(map[int]int, m)
	tasks[1] = 3
	tasks[2] = 5
	tasks[3] = 7
	tasks[4] = 10
	tasks[5] = 5
	tasks[6] = 100
	tasks[9] = 2

	totalPower := 0
	maxTime := 9

	for curTime := 1; curTime <= maxTime; curTime++ {
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
	// 3*2 + 5*3 + 7*4 + 10*2 + 5*6 + 2*3 = 105
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
