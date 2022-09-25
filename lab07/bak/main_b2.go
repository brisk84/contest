package main

import (
	"fmt"
	"strconv"
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
	// in := bufio.NewReader(os.Stdin)
	total := 1
	// fmt.Fscan(in, &total)

	for i := 0; i < total; i++ {
		n := 2
		m := 7
		// fmt.Fscan(in, &n, &m)

		hexs := make(map[rune]int, n*m) // ???
		// row := make([]rune, n)
		// mMap := make([][]rune, n*m)
		mMap := make([][]rune, m)

		// inp := []string{
		// 	"R.R.R.G",
		// 	".Y.G.G.",
		// 	"B.Y.V.V",
		// }
		// inp := []string{}
		// for nn := 0; nn < n; nn++ {
		// 	str := ""
		// 	fmt.Fscan(in, &str)
		// 	inp = append(inp, str)
		// }

		inp := []string{
			"G.B.R.G",
			".G.G.G.",
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
				fmt.Println(strconv.QuoteRune(h), h, count)

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

	h := (*field)[x][y]
	(*field)[x][y] = 0
	count++
	(*hexs)[h]--
	if (*hexs)[h] == 0 {
		return count
	}

	cy := y
	cx := x
	if y+2 < m {
		if (*field)[x][y+2] == h {
			cy += 2
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	if y-2 >= 0 {
		if (*field)[x][y-2] == h {
			cy -= 2
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	if y+1 < n && x+1 < n {
		if (*field)[x+1][y+1] == h {
			cx++
			cy++
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	if y-1 >= 0 && x+1 < n {
		if (*field)[x+1][y-1] == h {
			cx++
			cy--
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	if y-1 >= 0 && x-1 >= 0 {
		if (*field)[x-1][y-1] == h {
			cx--
			cy--
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	if y+1 < m && x-1 >= 0 {
		if (*field)[x-1][y+1] == h {
			cx--
			cy++
			count += GetRegCount(field, n, m, cx, cy, hexs)
		}
	}
	return count
}
