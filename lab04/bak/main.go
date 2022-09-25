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
		rep := make([]int, count)
		for j := 0; j < count; j++ {
			fmt.Fscan(in, &rep[j])
		}
		// rep := []int{1, 2, 3, 4, 5} // YES
		// rep := []int{1, 2, 3, 1}
		// rep := []int{2, 3, 4, 8, 5, 5, 5, 5} // YES
		// rep := []int{1, 1, 3, 2, 2} // YES
		// rep := []int{1, 1, 2, 3, 2} // NO
		valid := "NO"
		br := false
		for k, v := range rep {
			prevDay := k
			for kk, vv := range rep {
				subAbs := kk - prevDay
				if v == vv && subAbs < 2 {
					prevDay = kk
				}
				if v == vv && k != kk && prevDay != kk {
					br = true
					break
				}
			}
			if br {
				break
			}
		}
		if !br {
			valid = "YES"
		}
		s += fmt.Sprintln(valid)
	}
	fmt.Println(s)
}
