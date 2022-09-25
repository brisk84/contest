package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	total := 0
	fmt.Fscan(in, &total)

	// total := 1
	s := ""
	for i := 0; i < total; i++ {
		count := 0
		fmt.Fscan(in, &count)

		// rep := []int{1, 2, 3, 4, 5} // YES
		// rep := []int{1, 2, 3, 1} // NO
		// rep := []int{2, 3, 4, 8, 5, 5, 5, 5} // YES
		// rep := []int{1, 1, 3, 2, 2} // YES
		// rep := []int{1, 1, 2, 3, 2} // NO
		// count := len(rep)

		m := make(map[int][]int, count)
		for j := 0; j < count; j++ {
			taskNum := 0
			fmt.Fscan(in, &taskNum)
			m[taskNum] = append(m[taskNum], j)
		}

		valid := "YES"
		br := false
		for _, v := range m {
			prevDay := v[0]
			for _, vv := range v {
				sub := vv - prevDay
				if sub > 1 {
					valid = "NO"
					br = true
				}
				prevDay = vv
			}
			if br {
				break
			}
		}
		s += fmt.Sprintln(valid)
	}
	fmt.Println(s)
}
