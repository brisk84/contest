package main

import (
	"bufio"
	"fmt"
	"os"
)

var flDebug = false

type Hexa struct {
	LeftUpper  *Hexa
	Left       *Hexa
	LeftDown   *Hexa
	RightDown  *Hexa
	Right      *Hexa
	RightUpper *Hexa
	Color      rune
}

type Field = [][]rune

func main() {

	// s := ""
	in := bufio.NewReader(os.Stdin)
	total := 1
	fmt.Fscan(in, &total)

	for i := 0; i < total; i++ {
		n := 2
		m := 5
		fmt.Fscan(in, &n, &m)

		hexs := make(map[rune]int, n*m) // ???
		// row := make([]rune, n)
		// mMap := make([][]rune, n*m)
		mMap := make([][]rune, n*m)

		// inp := []string{
		// 	"A.J.J",
		// 	".J.I.",
		// }
		inp := []string{}
		for nn := 0; nn < n; nn++ {
			str := ""
			fmt.Fscan(in, &str)
			inp = append(inp, str)
		}

		// gx := 0
		// gy := 0

		for j := 0; j < n; j++ {
			mMap[j] = make([]rune, m)
			for k, v := range inp[j] {
				mMap[j][k] = v
				if v != '.' {
					hexs[v]++
				}
			}
		}

		no := false
		for gx := 0; gx < n; gx++ {
			for gy := 0; gy < m; gy++ {
				// if (gx + 1) % 2 == 0 {

				// }
				h := mMap[gx][gy]
				if h == '.' || h == 0 {
					continue
				}
				count := GetRegCount(&mMap, n, m, gx, gy, &hexs)
				// fmt.Println(h, count)
				_ = count

				if hexs[h] > 0 {
					no = true
					break
				}
			}
			if no {
				break
			}
		}
		if no {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

		// fmt.Println(hexs)
		// fmt.Println(inp)
		// fmt.Println(mMap)
	}
}

func GetRegCount(field *[][]rune, n int, m int, x int, y int, hexs *map[rune]int) int {
	count := 0
	// fmt.Println(x, y)
	h := (*field)[x][y]
	// fmt.Println(x, y, h, len(*hexs))
	// if h == 0 {
	// 	return 0
	// }
	(*field)[x][y] = 0
	count++
	(*hexs)[h]--
	if (*hexs)[h] == 0 {
		return count
	}

	cx := x
	cy := y
	if cy+2 < m {
		if (*field)[cx][cy+2] == h {
			cy += 2
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	cx = x
	cy = y
	if cy-2 >= 0 {
		if (*field)[cx][cy-2] == h {
			cy -= 2
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	cx = x
	cy = y
	if cy+1 < m && cx+1 < n {
		if (*field)[cx+1][cy+1] == h {
			cx++
			cy++
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	cx = x
	cy = y
	if cy-1 >= 0 && cx+1 < n {
		xxx := (*field)[cx+1][cy-1]
		if xxx == h {
			_ = xxx
		}
		if (*field)[cx+1][cy-1] == h {
			cx++
			cy--
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	cx = x
	cy = y
	if cy-1 >= 0 && cx-1 >= 0 {
		if (*field)[cx-1][cy-1] == h {
			cx--
			cy--
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	cx = x
	cy = y
	if cy+1 < m && cx-1 >= 0 {
		if (*field)[cx-1][cy+1] == h {
			cx--
			cy++
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	return count
}
