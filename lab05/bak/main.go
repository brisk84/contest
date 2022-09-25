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

var flDebug = true

func main() {
	in := bufio.NewReader(os.Stdin)
	total := 0
	if !flDebug {
		fmt.Fscan(in, &total)
	} else {
		total = 1
	}
	s := ""
	for i := 0; i < total; i++ {
		count := 0
		if !flDebug {
			fmt.Fscan(in, &count)
		}
		// rep := []int{1, 2, 3, 4, 5} // YES
		// rep := []int{1, 2, 3, 1} // NO
		// rep := []int{2, 3, 4, 8, 5, 5, 5, 5} // YES
		// rep := []int{1, 1, 3, 2, 2} // YES
		// rep := []int{1, 1, 2, 3, 2} // NO
		// count := len(rep)

		busy := [86400]bool{}

		// inps := []string{"23:59:59-23:59:59", "00:00:00-23:59:58"}
		// inps := []string{"24:00:00-23:59:59"}

		// inps := []string{"23:59:58-23:59:59", "00:00:00-23:59:58"}
		inps := []string{
			"18:12:49-18:22:00",
			"10:04:37-10:15:21",
			"22:18:53-22:39:39",
			"16:29:56-16:31:34",
			"14:18:22-14:28:15",
			"11:30:57-11:57:13",
			"02:42:35-02:48:36",
			"03:38:22-04:21:26",
			"16:07:27-16:26:43",
			"07:58:10-08:00:21",
			"00:33:25-00:35:36",
			"01:18:27-02:32:19",
			"18:22:19-18:40:13",
			"06:21:43-06:59:50",
			"07:12:36-07:14:43",
			"01:10:54-01:16:05",
			"05:57:44-06:20:29",
			"04:55:02-05:02:36",
			"17:45:01-18:06:50",
			"23:20:45-23:57:49",
			"05:11:24-05:11:24",
			"21:30:47-21:42:10",
			"14:38:54-14:40:59",
			"13:48:01-13:48:01",
			"11:14:32-23:43:48",
			"08:04:16-08:16:06",
			"19:11:13-19:22:34",
			"11:24:23-11:26:42",
			"17:00:18-17:07:22",
			"13:41:04-13:41:16",
			"12:22:21-12:51:26",
			"22:49:40-23:10:15",
			"04:29:27-04:46:34",
			"14:42:36-14:42:50",
			"14:52:57-15:49:08",
			"20:49:08-21:25:08",
			"09:12:28-09:16:07",
			"10:49:23-11:22:40",
			"18:45:41-18:59:04",
			"00:17:48-00:22:34",
		}
		count = len(inps)
		// inps := make([]string, count)

		if !flDebug {
			for j := 0; j < count; j++ {
				fmt.Fscan(in, &inps[j])
			}
		}
		br := false
		// count := len(inps)
		for _, v := range inps {
			// inp := "02:46:00-03:14:59"
			valid, s1, s2 := ValidateTime(v)
			// fmt.Println(s1, s2)
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
