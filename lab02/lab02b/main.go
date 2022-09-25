package main

import (
	"fmt"
)

func main() {
	total := 0
	fmt.Scan(&total)
	goods := make([]map[int]int, total)

	for i := 0; i < total; i++ {
		count := 0
		fmt.Scan(&count)
		prices := make([]int, count+1)
		for j := 0; j <= count; j++ {
			price := 0
			fmt.Scanf("%d", &price)
			prices[j] = price
		}
		goods[i] = make(map[int]int, count)
		for _, v := range prices {
			if _, ok := goods[i][v]; !ok {
				goods[i][v] = 0
			}
			goods[i][v]++
		}
	}
	for _, v := range goods {
		sum := 0
		for k, vv := range v {
			sum += k*(vv/3)*2 + k*(vv%3)
		}
		fmt.Println(sum)
	}
}
