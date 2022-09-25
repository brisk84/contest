package main

import "fmt"

type Num struct {
	A int
	B int
}

func main() {
	count := 0
	nums := []Num{}
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		num := Num{}
		fmt.Scan(&num.A, &num.B)
		nums = append(nums, num)
	}
	for _, v := range nums {
		fmt.Println(v.A + v.B)
	}
}
