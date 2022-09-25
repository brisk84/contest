package main

import (
	"fmt"
)

func getMaxReqs(in string) string {
	ret := ""

	return ret
}

func getNext(row, col, rows, cols int) (int, int) {
	if row >= rows {
		if col >= cols {
			return 0, 0
		}
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

func main() {

	// field := []string{
	// 	".*....",
	// 	".*.***",
	// 	".***.*",
	// 	".....*",
	// 	"......",
	// }

	field := [][]byte{
		{'.', '*', '.', '.', '.', '.'},
		{'.', '*', '.', '*', '*', '*'},
		{'.', '*', '*', '*', '.', '*'},
		{'.', '.', '.', '.', '.', '*'},
		{'.', '*', '.', '.', '.', '.'},
	}

	rows := 5
	cols := 6

	row := 0
	col := 0

	row, col = findFirst(field, rows, cols)
	// field[row][col] = '.'
	// fmt.Println(row, col)

	path := ""

	for {
		dir := ""
		field[row][col] = '.'
		row, col, dir = findNext(field, rows, cols, row, col)
		if row == -1 {
			break
		}
		path += dir
		// fmt.Println(path)
	}
	fmt.Println(path)

	// ret := getMaxReqs("100100100100100100100")
	// fmt.Println(ret)

	// return

	// in := bufio.NewReader(os.Stdin)
	// total := 0
	// fmt.Fscan(in, &total)

	// s := ""
	// for i := 0; i < total; i++ {
	// 	inp := ""
	// 	fmt.Fscan(in, &inp)
	// 	s += fmt.Sprintln(decode(inp))
	// }
	// fmt.Println(s)
}
