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

	var minTime uint64
	var maxTime uint64
	tasks := make(map[uint64]uint64, m)
	for i := 0; i < m; i++ {
		var startTime uint64
		var workTime uint64
		fmt.Fscan(in, &startTime, &workTime)
		tasks[startTime] = workTime
		if i == 0 {
			minTime = startTime
		}
		if i == m-1 {
			maxTime = startTime
		}
	}

	var totalPower uint64
	for curTime := minTime; curTime <= maxTime; curTime++ {
		if curTask, ok := tasks[curTime]; ok {
			for k := range procs {
				if procs[k].BusyUntil <= curTime {
					procs[k].BusyUntil = curTime + curTask
					totalPower += curTask * procs[k].Power
					break
				}
			}
		}
	}
	fmt.Println(totalPower)
}
