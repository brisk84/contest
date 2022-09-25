package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ValidateTime(inp string) (bool, int, int) {
	hh1, _ := strconv.Atoi(inp[0:2])
	mm1, _ := strconv.Atoi(inp[3:5])
	ss1, _ := strconv.Atoi(inp[6:8])
	hh2, _ := strconv.Atoi(inp[9:11])
	mm2, _ := strconv.Atoi(inp[12:14])
	ss2, _ := strconv.Atoi(inp[15:17])
	if hh1 < 0 || hh1 > 23 {
		return false, 0, 0
	}
	if hh2 < 0 || hh2 > 23 {
		return false, 0, 0
	}
	if mm1 < 0 || mm1 > 59 {
		return false, 0, 0
	}
	if mm2 < 0 || mm2 > 59 {
		return false, 0, 0
	}
	if ss1 < 0 || ss1 > 59 {
		return false, 0, 0
	}
	if ss2 < 0 || ss2 > 59 {
		return false, 0, 0
	}

	stm1 := inp[0:8]
	stm2 := inp[9:17]
	tm1, err := time.Parse("15:04:05", stm1)
	if err != nil {
		return false, 0, 0
	}
	tm2, err := time.Parse("15:04:05", stm2)
	if err != nil {
		return false, 0, 0
	}
	if tm1.After(tm2) {
		return false, 0, 0
	}
	// fmt.Println(tm1, tm2)
	s1 := ss1 + mm1*60 + hh1*60*60
	s2 := ss2 + mm2*60 + hh2*60*60

	return true, s1, s2
}

func main() {
	in := bufio.NewReader(os.Stdin)
	total := 0
	fmt.Fscan(in, &total)

	// total := 1
	s := ""
	for i := 0; i < total; i++ {
		count := 0
		fmt.Fscan(in, &count)

		busy := [86400]bool{}

		// inps := []string{"23:59:59-23:59:59", "00:00:00-23:59:58"}
		// inps := []string{"24:00:00-23:59:59"}
		inps := make([]string, count)
		for j := 0; j < count; j++ {
			fmt.Fscan(in, &inps[j])
		}
		br := false
		// count := len(inps)
		for _, v := range inps {
			// inp := "02:46:00-03:14:59"
			valid, s1, s2 := ValidateTime(v)
			if !valid {
				s += fmt.Sprintln("NO")
				br = true
				break
			}
			for i := s1; i <= s2; i++ {
				if busy[i] {
					s += fmt.Sprintln("NO")
					br = true
					break
				}
				busy[i] = true
			}
			if br {
				break
			}
		}
		if !br {
			s += fmt.Sprintln("YES")
		}

	}
	fmt.Println(s)
}
