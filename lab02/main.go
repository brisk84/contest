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
	s := ""
	for i := 0; i < total; i++ {
		count := 0
		fmt.Fscan(in, &count)
		prices := make(map[int]int, count+1)
		for j := 0; j < count; j++ {
			price := 0
			fmt.Fscan(in, &price)
			prices[price]++
		}
		goods := 0
		for k, vv := range prices {
			goods += k*(vv/3)*2 + k*(vv%3)
		}
		s += fmt.Sprintln(goods)
	}
	fmt.Print(s)
}
