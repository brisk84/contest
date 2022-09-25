package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dev struct {
	Num int
	Lev int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	total := 0
	fmt.Fscan(in, &total)

	// total := 1
	// s := ""
	for i := 0; i < total; i++ {
		count := 0
		fmt.Fscan(in, &count)
		devsInt := make([]int, count)
		for j := 0; j < count; j++ {
			fmt.Fscan(in, &devsInt[j])
		}
		// devsInt := []int{2, 1, 3, 1, 1, 4}
		// devsInt := []int{5, 4, 8, 2, 1, 1, 4, 4}

		devs := make([]Dev, len(devsInt))
		for k, v := range devsInt {
			dev := Dev{}
			dev.Num = k + 1
			dev.Lev = v
			devs[k] = dev
		}

		for len(devs) >= 2 {
			first := devs[0].Num
			firstDev := devs[0]
			firstLev := devs[0].Lev
			secondDev := Dev{}
			second := 0
			min := 100
			minPos := 100
			devs = RemoveDev(devs, firstDev)
			for k, v := range devs {
				sub := firstLev - v.Lev
				if sub < 0 {
					sub = 0 - sub
				}
				if sub < min {
					min = sub
					second = v.Num
					secondDev = v
					minPos = k
				} else if sub == min {
					if k <= minPos {
						min = sub
						second = v.Num
						secondDev = v
						minPos = k
					}
				}
			}
			devs = RemoveDev(devs, secondDev)
			fmt.Println(first, second)
		}
		fmt.Println()
	}
}

func RemoveDev(devs []Dev, dev Dev) []Dev {
	newDevs := []Dev{}
	for _, v := range devs {
		if v.Num != dev.Num {
			newDevs = append(newDevs, v)
		}
	}
	return newDevs
}
