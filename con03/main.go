package main

import (
	"bufio"
	"fmt"
	"os"
)

func decode(in string) string {
	ret := ""

	for len(in) > 0 {
		s2 := in[0:2]
		if s2 == "00" {
			ret += "a"
			in = in[2:]
			continue
		}
		if s2 == "11" {
			ret += "d"
			in = in[2:]
			continue
		}
		s3 := in[0:3]
		if s3 == "100" {
			ret += "b"
			in = in[3:]
			continue
		}
		if s3 == "101" {
			ret += "c"
			in = in[3:]
			continue
		}
	}
	return ret
}

func main() {
	// ret := decode("100100100100100100100")
	// fmt.Println(ret)

	// return

	in := bufio.NewReader(os.Stdin)
	total := 0
	fmt.Fscan(in, &total)

	s := ""
	for i := 0; i < total; i++ {
		inp := ""
		fmt.Fscan(in, &inp)
		s += fmt.Sprintln(decode(inp))
	}
	fmt.Println(s)
}
