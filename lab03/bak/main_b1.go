package main

import (
	"fmt"
)

type Dev struct {
	Num int
	Lev int
}

func main() {
	// in := bufio.NewReader(os.Stdin)
	// total := 0
	// fmt.Fscan(in, &total)

	total := 1
	// s := ""
	for i := 0; i < total; i++ {
		// count := 0
		// fmt.Fscan(in, &count)
		// devsInt := make([]int, count)
		// for j := 0; j < count; j++ {
		// 	fmt.Fscan(in, &devsInt[j])
		// }
		// devsInt := []int{2, 1, 3, 1, 1, 4}
		devsInt := []int{5, 4, 8, 2, 1, 1, 4, 4}
		// 1 2
		// 3 7
		// 4 5
		// 6 8

		devs := make(map[int]Dev, len(devsInt))
		for k, v := range devsInt {
			dev := Dev{}
			dev.Num = k + 1
			dev.Lev = v
			devs[k] = dev
		}

		for len(devs) >= 2 {
			minKey := GetMinKey(&devs)
			first := devs[minKey].Num
			firstLev := devs[minKey].Lev
			second := 0
			minPos := 100
			min := 100
			delete(devs, first-1)
			for k, v := range devs {
				sub := firstLev - v.Lev
				if sub < 0 {
					sub = 0 - sub
				}
				if sub <= min && k < minPos {
					min = sub
					minPos = v.Num
					second = v.Num
				}
			}
			delete(devs, second-1)
			fmt.Println(first, second)
		}
		fmt.Println()
	}
}

func GetMinKey(m *map[int]Dev) int {
	min := 100
	for k, _ := range *m {
		if k < min {
			min = k
		}
	}
	return min
}

// func main() {
// 	// in := bufio.NewReader(os.Stdin)
// 	// total := 0
// 	// fmt.Fscan(in, &total)
// 	total := 1
// 	s := ""
// 	for i := 0; i < total; i++ {
// 		// count := 0
// 		// fmt.Fscan(in, &count)
// 		// count := 6
// 		// devs := make([]int, count)
// 		// for j := 0; j < count; j++ {
// 		// 	fmt.Fscan(in, &devs[j])
// 		// }
// 		devs := []int{2, 1, 3, 1, 1, 4}

// 		removed := 0
// 		for len(devs) >= 2 {
// 			first := 1 + removed
// 			firstLevel := devs[0]
// 			second := 0
// 			// secondLevel := 100
// 			devs = devs[1:]
// 			removed++
// 			min := 100

// 			for k, v := range devs {
// 				// sub := first - v
// 				sub := firstLevel - v
// 				if sub < 0 {
// 					sub = 0 - sub
// 				}

// 				if sub < min {
// 					min = sub
// 					second = k + 1 + removed
// 					// secondLevel = v
// 				}
// 			}
// 			fmt.Println(first, second)
// 			devs = append(devs[0:second-removed-1], devs[second-removed:]...)
// 			removed++

// 			s += fmt.Sprintf("%d %d\n", first, second)
// 		}
// 	}
// 	fmt.Print(s)
// }
