package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Proc struct {
	Power     uint64
	BusyUntil uint64
}

type Task struct {
	StartTime uint64
	WorkTime  uint64
}

func main() {

	in := bufio.NewReader(os.Stdin)
	n := 0 // procs
	m := 0 // tasks
	fmt.Fscan(in, &n, &m)
	procs := make([]Proc, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &procs[i].Power)
	}
	sort.Slice(procs, func(i, j int) bool {
		return procs[i].Power < procs[j].Power
	})

	// 	tasks := make(map[uint64]uint64, m)
	tasks := make([]Task, m)
	for i := 0; i < m; i++ {
		var startTime uint64
		var workTime uint64
		fmt.Fscan(in, &startTime, &workTime)
		// 		tasks[startTime] = workTime
		newTask := Task{
			StartTime: startTime,
			WorkTime:  workTime,
		}
		tasks = append(tasks, newTask)
	}

	var totalPower uint64
	var curTime uint64
	for _, curTask := range tasks {
		curTime = curTask.StartTime
		for k := range procs {
			if procs[k].BusyUntil <= curTime {
				procs[k].BusyUntil = curTime + curTask.WorkTime
				totalPower += curTask.WorkTime * procs[k].Power
				break
			}
		}
	}
	fmt.Println(totalPower)
}
