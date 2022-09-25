package main

import (
	"bufio"
	"fmt"
	"os"
)

func isLeap(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		return true
	}
	if year%400 == 0 {
		return true
	}
	return false
}

func isValid(day, month, year int) bool {
	if (month == 4 || month == 6 || month == 9 || month == 11) && (day > 30) {
		return false
	} else if month == 2 {
		if day > 29 {
			return false
		}
		if day == 29 && !isLeap(year) {
			return false
		}
	}
	return true
}

func main() {
	// ret := isValid(31, 11, 1999)
	// fmt.Println(ret)

	// return

	in := bufio.NewReader(os.Stdin)
	total := 0
	fmt.Fscan(in, &total)

	s := ""
	for i := 0; i < total; i++ {
		day := 0
		month := 0
		year := 0
		fmt.Fscan(in, &day, &month, &year)

		if isValid(day, month, year) {
			s += fmt.Sprintln("YES")
			continue
		}
		s += fmt.Sprintln("NO")
	}
	fmt.Println(s)
}
