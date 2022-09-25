package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Frame struct {
	StartRow int
	StartCol int
	EndRow   int
	EndCol   int
	Includes int
}

func getNext(row, col, rows, cols int) (int, int) {
	if row >= rows-1 {
		if col >= cols-1 {
			return -1, -1
		}
		return row, col + 1
	}
	if col >= cols-1 {
		return row + 1, 0
	}
	return row, col + 1
}

func findFirst(field *[][]byte, rows, cols int) (int, int) {
	row := 0
	col := 0
	for {
		if (*field)[row][col] != '*' {
			row, col = getNext(row, col, rows, cols)
			// fmt.Println(row, col)
			if row == -1 {
				return -1, -1
			}
			continue
		}
		return row, col
	}
}

func findNext(field *[][]byte, rows, cols, row, col int) (int, int) {
	if col > 0 && (*field)[row][col-1] == '*' {
		return row, col - 1
	}
	if col < cols-1 && (*field)[row][col+1] == '*' {
		return row, col + 1
	}
	if row > 0 && (*field)[row-1][col] == '*' {
		return row - 1, col
	}
	if row < rows-1 && (*field)[row+1][col] == '*' {
		return row + 1, col
	}
	return -1, -1
}

func findFrame(field *[][]byte, rows, cols int) *Frame {
	row, col := findFirst(field, rows, cols)
	if row == -1 {
		return nil
	}
	ret := Frame{}
	ret.StartCol = col
	ret.StartRow = row
	ret.EndCol = col
	ret.EndRow = row
	(*field)[row][col] = '.'

	for {
		row, col = findNext(field, rows, cols, row, col)
		if row == -1 {
			break
		}
		if ret.StartCol > col {
			ret.StartCol = col
		}
		if ret.EndCol < col {
			ret.EndCol = col
		}
		if ret.StartRow > row {
			ret.StartRow = row
		}
		if ret.EndRow < row {
			ret.EndRow = row
		}
		(*field)[row][col] = '.'
	}

	return &ret
}

func loadFrames(field *[][]byte, rows, cols int) []Frame {
	frames := []Frame{}
	frame := findFrame(field, rows, cols)

	// cnt := 1
	for frame != nil {
		frames = append(frames, *frame)
		frame = findFrame(field, rows, cols)
		// fmt.Println("cnt:", cnt)
		// cnt++
	}
	// fmt.Println(len(frames))
	// frames = append(frames, *frame)
	return frames
}

func getIncludes(frames []Frame) []int {
	for k, v := range frames {
		for kk, vv := range frames {
			if k == kk {
				continue
			}
			if vv.StartRow < v.StartRow {
				continue
			}
			if vv.StartCol < v.StartCol {
				continue
			}
			if vv.StartRow > v.EndRow {
				continue
			}
			if vv.StartCol > v.EndCol {
				continue
			}
			frames[kk].Includes++
			// if vv.StartRow > v.StartRow && vv.StartCol > v.StartCol && vv.StartRow < v.EndRow && vv.StartCol < v.EndCol {
			// 	frames[kk].Includes++
			// }
		}
	}
	ret := make([]int, len(frames))
	for k, v := range frames {
		ret[k] = v.Includes
	}
	return ret
}

func getIncludes2(field *[][]byte, rows, cols int) string {
	frames := loadFrames(field, rows, cols)
	for k, v := range frames {
		for kk, vv := range frames {
			if k == kk {
				continue
			}
			if vv.StartRow < v.StartRow {
				continue
			}
			if vv.StartCol < v.StartCol {
				continue
			}
			if vv.StartRow > v.EndRow {
				continue
			}
			if vv.StartCol > v.EndCol {
				continue
			}
			frames[kk].Includes++
			// if vv.StartRow > v.StartRow && vv.StartCol > v.StartCol && vv.StartRow < v.EndRow && vv.StartCol < v.EndCol {
			// 	frames[kk].Includes++
			// }
		}
	}
	// ret := make([]int, len(frames))
	// for k, v := range frames {
	// 	ret[k] = v.Includes
	// }
	// return ret

	sort.SliceStable(frames, func(i, j int) bool {
		return frames[i].Includes < frames[j].Includes
	})
	s := ""
	for _, v := range frames {
		s += fmt.Sprintf("%d ", v.Includes)
	}
	return s
	// ss := fmt.Sprint(includes)
	// fmt.Println(ss[1 : len(ss)-1])
}

func main() {

	in := bufio.NewReader(os.Stdin)

	// f, err := os.Open("/Users/brisk/vscode/contest/con08/in.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// in := bufio.NewReader(f)

	total := 0
	fmt.Fscan(in, &total)

	// s := ""
	// cntr := 0
	for i := 0; i < total; i++ {
		rows := 0
		cols := 0
		fmt.Fscan(in, &rows, &cols)

		field := [][]byte{}
		for curRow := 0; curRow < rows; curRow++ {
			fRow := make([]byte, cols)
			fmt.Fscan(in, &fRow)
			field = append(field, fRow)
		}

		// frames := loadFrames(&field, rows, cols)
		// includes := getIncludes(frames)

		// frames := loadFrames(&field, rows, cols)
		includes := getIncludes2(&field, rows, cols)
		fmt.Println(includes)

		// sort.SliceStable(includes, func(i, j int) bool {
		// 	return includes[i] < includes[j]
		// })
		// ss := fmt.Sprint(includes)
		// fmt.Println(ss[1 : len(ss)-1])

		// cntr++
		// if cntr%10 == 0 {
		// fmt.Print(s)
		// s = ""
		// }

		// s := ""
		// for _, v := range includes {
		// 	s += fmt.Sprintf("%d ", v)
		// }
		// fmt.Println(s)

	}
	// fmt.Print(s)
}
