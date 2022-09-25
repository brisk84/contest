package main

import (
	"bufio"
	"fmt"
	"os"
)

func getNext(row, col, rows, cols int) (int, int) {
	if col >= cols-1 {
		if row >= rows-1 {
			return 0, 0
		}
		return row + 1, 0
	}
	return row, col + 1

}

func getPathsCount(field [][]byte, row, col, rows, cols int) int {
	count := 0
	if col > 0 && field[row][col-1] == '*' {
		count++
	}
	if col < cols-1 && field[row][col+1] == '*' {
		count++
	}
	if row > 0 && field[row-1][col] == '*' {
		count++
	}
	if row < rows-1 && field[row+1][col] == '*' {
		count++
	}
	return count
}

func findFirst(field [][]byte, rows, cols int) (int, int) {
	row := 0
	col := 0
	for {
		if field[row][col] != '*' {
			row, col = getNext(row, col, rows, cols)
			continue
		}
		if getPathsCount(field, row, col, rows, cols) > 1 {
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

	in := bufio.NewReader(os.Stdin)

	// f, err := os.Open("/Users/brisk/vscode/contest/con05/in.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// in := bufio.NewReader(f)

	total := 0
	fmt.Fscan(in, &total)

	// s := ""
	for i := 0; i < total; i++ {
		rows := 0
		cols := 0
		fmt.Fscan(in, &rows, &cols)

		field := [][]byte{}
		for curRow := 0; curRow < rows; curRow++ {
			sRow := ""
			fmt.Fscan(in, &sRow)
			fRow := make([]byte, 100)
			for k, v := range sRow {
				fRow[k] = byte(v)
			}
			field = append(field, fRow)
		}

		row, col := findFirst(field, rows, cols)

		path := ""
		for {
			dir := ""
			field[row][col] = '.'
			row, col, dir = findNext(field, rows, cols, row, col)
			if row == -1 {
				break
			}
			path += dir
		}
		fmt.Println(path)
	}
}
