package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Proc struct {
	Power     uint64
	BusyUntil uint64
}

func main() {
	f, err := os.Open("/Users/brisk/vscode/contest/lab08/tests/07")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)
	// in := bufio.NewReader(os.Stdin)
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

	fw, err := os.Create("/Users/brisk/vscode/contest/lab08/log.txt")
	defer fw.Close()
	out := bufio.NewWriter(fw)

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
					if totalPower == 151181876613406 {
						xx := 0
						_ = xx
					}
					procs[k].BusyUntil = curTime + curTask
					totalPower += curTask * procs[k].Power
					fmt.Fprintln(out, totalPower)
					break
				}
			}
		}
	}
	fmt.Println(totalPower)
}
