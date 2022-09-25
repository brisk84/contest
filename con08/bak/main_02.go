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

func findFirst(field [][]byte, rows, cols int) (int, int) {
	row := 0
	col := 0
	for {
		if field[row][col] != '*' {
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

func findNext(field [][]byte, rows, cols, row, col int) (int, int, string) {
	if col > 0 && field[row][col-1] == '*' {
		return row, col - 1, "L"
	}
	if col < cols-1 && field[row][col+1] == '*' {
		return row, col + 1, "R"
	}
	if row > 0 && field[row-1][col] == '*' {
		return row - 1, col, "U"
	}
	if row < rows-1 && field[row+1][col] == '*' {
		return row + 1, col, "D"
	}
	return -1, -1, ""
}

func findFrame(field [][]byte, rows, cols int) *Frame {
	row, col := findFirst(field, rows, cols)
	if row == -1 {
		return nil
	}
	ret := Frame{}
	ret.StartCol = col
	ret.StartRow = row
	ret.EndCol = col
	ret.EndRow = row
	field[row][col] = '.'

	for {
		row, col, _ = findNext(field, rows, cols, row, col)
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
		field[row][col] = '.'
	}

	return &ret
}

func loadFrames(field [][]byte, rows, cols int) []Frame {
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
			if vv.StartRow > v.StartRow && vv.StartCol > v.StartCol && vv.StartRow < v.EndRow && vv.StartCol < v.EndCol {
				frames[kk].Includes++
			}
		}
	}
	ret := make([]int, len(frames))
	for k, v := range frames {
		ret[k] = v.Includes
	}
	return ret
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

	s := ""
	for i := 0; i < total; i++ {
		rows := 0
		cols := 0
		fmt.Fscan(in, &rows, &cols)

		field := [][]byte{}
		for curRow := 0; curRow < rows; curRow++ {
			// sRow := ""
			// fmt.Fscan(in, &sRow)
			fRow := make([]byte, cols)
			// for k, v := range sRow {
			// 	fRow[k] = byte(v)
			// }
			fmt.Fscan(in, &fRow)

			field = append(field, fRow)
			// fmt.Println(fRow)
		}
		// fmt.Println(field)

		frames := loadFrames(field, rows, cols)
		includes := getIncludes(frames)
		sort.Slice(includes, func(i, j int) bool {
			return includes[i] < includes[j]
		})
		ss := fmt.Sprint(includes)
		s += fmt.Sprintln(ss[1 : len(ss)-1])

		// s := ""
		// for _, v := range includes {
		// 	s += fmt.Sprintf("%d ", v)
		// }
		// fmt.Println(s)

	}
	fmt.Print(s)
}
